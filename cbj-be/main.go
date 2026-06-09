package main

import (
	"cbj-be/controller/encode"
	"cbj-be/controller/product"
	"cbj-be/internal/signer"
	"cbj-be/models"
	"cbj-be/router"
	"cbj-be/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// default router
	r := gin.Default()
	root := r.Group("cbj-launchpad")

	// initialization
	config, err := utils.LoadConfig("./conf/app.ini")
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

	productController := product.NewProductController(db)
	encodeController := encode.NewEncodeController(sgn)
	allocationController := product.NewAllocationController()

	router.ApiRouterInit(root, productController, encodeController, allocationController)
	router.AdminRouterInit(root)

	http.ListenAndServe(":8080", r)
	r.Run()
}
