package main

import (
	"Backend_WebProg_RentalCarManagement/controller"
	"Backend_WebProg_RentalCarManagement/database"
	"Backend_WebProg_RentalCarManagement/initializer"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.EnvLoader()
	database.ConnectDatabase()
}

func main() {
	r := gin.Default()

	// Cars Routes
	r.GET("/api/cars", controller.CarsIndex)
	r.GET("/api/cars/:id", controller.CarsShow)
	r.POST("/api/cars", controller.CarsCreate)
	r.PUT("/api/cars/:id", controller.CarsUpdate)
	r.DELETE("/api/cars/", controller.CarsDelete)

	// Maintenance History Routes
	r.GET("/api/maintenance", controller.MaintenanceHistoryIndex)
	r.GET("/api/maintenance/:id", controller.MaintenanceHistoryShow)
	r.POST("/api/maintenance", controller.MaintenanceHistoryCreate)
	r.PUT("/api/maintenance/:id", controller.MaintenanceHistoryUpdate)
	r.DELETE("/api/maintenance", controller.MaintenanceHistoryDelete)
	r.GET("/api/maintenance/car/:id", controller.MaintenanceHistoryShowByCarID)

	// Payments Routes
	r.GET("/api/payments", controller.PaymentIndex)
	r.GET("/api/payments/:id", controller.PaymentShow)
	r.POST("/api/payments", controller.PaymentCreate)
	r.PUT("/api/payments/:id", controller.PaymentUpdate)
	r.DELETE("/api/payments", controller.PaymentDelete)

	// Rentals Routes
	r.GET("/api/rentals", controller.RentalIndex)
	r.GET("/api/rentals/:id", controller.RentalShow)
	r.POST("/api/rentals", controller.RentalCreate)
	r.PUT("/api/rentals/:id", controller.RentalUpdate)
	r.DELETE("/api/rentals", controller.RentalDelete)

	// Taxes Routes
	r.GET("/api/taxes", controller.TaxIndex)
	r.GET("/api/taxes/:id", controller.TaxShow)
	r.POST("/api/taxes", controller.TaxCreate)
	r.PUT("/api/taxes/:id", controller.TaxUpdate)
	r.DELETE("/api/taxes", controller.TaxDelete)
	r.GET("/api/taxes/car/:id", controller.GetTaxesByCarID)

	r.Run(os.Getenv("PORT"))
}
