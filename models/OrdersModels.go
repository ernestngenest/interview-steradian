package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CarID           uint      `json:"car_id" gorm:"not null"`
	OrderDate       time.Time `json:"order_date" gorm:"type:date;not null"`
	PickupDate      time.Time `json:"pickup_date" gorm:"type:date;not null"`
	DropoffDate     time.Time `json:"dropoff_date" gorm:"type:date;not null"`
	PickupLocation  string    `json:"pickup_location" gorm:"type:varchar(50);not null"`
	DropoffLocation string    `json:"dropoff_location" gorm:"type:varchar(50);not null"`
	Car             Car       `json:"-" gorm:"foreignKey:CarID"`
}

type OrderRequest struct {
	CarID           uint   `json:"car_id"`
	OrderDate       string `json:"order_date"`
	PickupDate      string `json:"pickup_date"`
	DropoffDate     string `json:"dropoff_date"`
	PickupLocation  string `json:"pickup_location"`
	DropoffLocation string `json:"dropoff_location"`
}
