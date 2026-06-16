package main

import (
	"cbj-be/controller/encode"
	"cbj-be/controller/product"
	"cbj-be/internal/signer"
	"cbj-be/models"
	"cbj-be/onchain"
	"cbj-be/router"
	"cbj-be/utils"
	"context"
	"fmt"
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	// default router
	r := gin.Default()
	root := r.Group("cbj-launchpad")

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	// initialization
	configPath := fmt.Sprintf("./conf/app-%s.ini", env)
	config, err := utils.LoadConfig(configPath)
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	db, err := models.NewDB(&config.MySQL)
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	sgn, err := signer.NewSigner(config.Signer.PrivateKey)
	if err != nil {
		panic("failed to create signer: " + err.Error())
	}

	startListener(config, db)

	productController := product.NewProductController(db)
	encodeController := encode.NewEncodeController(sgn)
	allocationController := product.NewAllocationController()
	registerController := product.NewRegisterController(db)

	router.ApiRouterInit(root, productController, encodeController, allocationController, registerController)
	router.AdminRouterInit(root)

	http.ListenAndServe(":8080", r)
	r.Run()
}

func startListener(config *utils.Config, db *gorm.DB) {
	listener, err := onchain.NewListener(config.OnChainParameters.WsURL, db)
	if err != nil {
		panic("failed to create listener: " + err.Error())
	}
	if err := listener.LoadSalesFromDB(); err != nil {
		panic("load sales failed: " + err.Error())
	}

	listener.RegisterHandler(onchain.NewRegisterHandler(listener))
	listener.RegisterHandler(onchain.NewParticipateHandler(listener))

	go listener.Start(context.Background())
}
