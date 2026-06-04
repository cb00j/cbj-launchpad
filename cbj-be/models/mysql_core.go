package models

//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
import (
	"cbj-be/utils"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(c *utils.MySQLConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.IP, c.Port, c.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true, // print long SQL with field names
		//SkipDefaultTransaction: true, // disable default transactions
	})
	if err != nil {
		return nil, fmt.Errorf("open mysql: %w", err)
	}
	return db, nil
}
