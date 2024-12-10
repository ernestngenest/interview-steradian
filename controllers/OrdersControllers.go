package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"main.go/configs"
	"main.go/models"
)

// GetOrders all
func GetOrders(c *gin.Context) {
	var orders []models.Order
	result := configs.DB.Preload("Car").Find(&orders)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch orders",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": orders,
	})
}

// GetOrder search by id
func GetOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	result := configs.DB.Preload("Car").First(&order, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Order not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}

func CreateOrder(c *gin.Context) {
	var orderReq models.OrderRequest
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Parse dates
	layout := "2006-01-02"
	pickupDate, err := time.Parse(layout, orderReq.PickupDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid pickup date format. Use YYYY-MM-DD",
		})
		return
	}

	dropoffDate, err := time.Parse(layout, orderReq.DropoffDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid dropoff date format. Use YYYY-MM-DD",
		})
		return
	}

	var orderDate time.Time
	if orderReq.OrderDate != "" {
		orderDate, err = time.Parse(layout, orderReq.OrderDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid order date format. Use YYYY-MM-DD",
			})
			return
		}
	} else {
		orderDate = time.Now()
	}

	order := models.Order{
		CarID:           orderReq.CarID,
		OrderDate:       orderDate,
		PickupDate:      pickupDate,
		DropoffDate:     dropoffDate,
		PickupLocation:  orderReq.PickupLocation,
		DropoffLocation: orderReq.DropoffLocation,
	}

	// validasi mobil
	var car models.Car
	if err := configs.DB.First(&car, order.CarID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid car ID",
		})
		return
	}

	// Validate dates
	if order.PickupDate.After(order.DropoffDate) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Pickup date must be before dropoff date",
		})
		return
	}

	result := configs.DB.Create(&order)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create order",
		})
		return
	}

	// fetch order yang di create beserta detail car
	configs.DB.Preload("Car").First(&order, order.ID)

	c.JSON(http.StatusCreated, gin.H{
		"data": order,
	})
}

func UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	// validasi
	if err := configs.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Order not found",
		})
		return
	}

	// masukan date original
	originalOrderDate := order.OrderDate

	// update data setelah validasi
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// verifikasi car yg di update exist
	if order.CarID != 0 {
		var car models.Car
		if err := configs.DB.First(&car, order.CarID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid car ID",
			})
			return
		}
	}

	order.OrderDate = originalOrderDate

	// Validate date
	if order.PickupDate.After(order.DropoffDate) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Pickup date must be before dropoff date",
		})
		return
	}

	// Save
	result := configs.DB.Save(&order)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update order",
		})
		return
	}

	// fetch car yg udh di update (sama kayak yg diatas)
	configs.DB.Preload("Car").First(&order, order.ID)

	c.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	// validasi
	if err := configs.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Order not found",
		})
		return
	}

	// Delete
	result := configs.DB.Delete(&order)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete order",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order deleted successfully",
	})
}
