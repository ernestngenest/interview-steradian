package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/configs"
	"main.go/models"
)

func GetCars(c *gin.Context) {
	var cars []models.Car
	result := configs.DB.Find(&cars)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch cars",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": cars,
	})
}

func GetCar(c *gin.Context) {
	id := c.Param("id")
	var car models.Car

	result := configs.DB.First(&car, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Car not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": car,
	})
}

func CreateCar(c *gin.Context) {
	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := configs.DB.Create(&car)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create car",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": car,
	})
}

func UpdateCar(c *gin.Context) {
	id := c.Param("id")
	var car models.Car

	if err := configs.DB.First(&car, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Car not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := configs.DB.Save(&car)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update car",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": car,
	})
}

func DeleteCar(c *gin.Context) {
	id := c.Param("id")
	var car models.Car

	if err := configs.DB.First(&car, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Car not found",
		})
		return
	}

	result := configs.DB.Delete(&car)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete car",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Car deleted successfully",
	})
}
