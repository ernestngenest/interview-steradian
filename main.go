package main

import (
	"github.com/gin-gonic/gin"
	"main.go/configs"
	"main.go/controllers"
	"main.go/models"
)

func init() {
	configs.LoadEnvVariables()
	configs.ConnectToDB()
	configs.DB.AutoMigrate(&models.Car{}, &models.Order{})
}

func main() {
	r := gin.Default()

	// Car routes
	cars := r.Group("/cars")
	{
		cars.GET("/", controllers.GetCars)
		cars.GET("/:id", controllers.GetCar)
		cars.POST("/", controllers.CreateCar)
		cars.PUT("/:id", controllers.UpdateCar)
		cars.DELETE("/:id", controllers.DeleteCar)
	}

	// Order routes
	orders := r.Group("/orders")
	{
		orders.GET("/", controllers.GetOrders)
		orders.GET("/:id", controllers.GetOrder)
		orders.POST("/", controllers.CreateOrder)
		orders.PUT("/:id", controllers.UpdateOrder)
		orders.DELETE("/:id", controllers.DeleteOrder)
	}

	r.Run()
}
