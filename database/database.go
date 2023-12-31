package database

import (
	"Backend_WebProg_RentalCarManagement/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open(os.Getenv("DB")))
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&models.Rental{})
	database.AutoMigrate(&models.MaintenanceHistory{})
	database.AutoMigrate(&models.Taxes{})
	database.AutoMigrate(&models.Payment{})
	database.AutoMigrate(&models.Car{})

	DB = database
}
