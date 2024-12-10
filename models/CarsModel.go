package models

import (
	"mime/multipart"

	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	CarName   string  `json:"car_name" gorm:"type:varchar(50);not null"`
	DayRate   float64 `json:"day_rate" gorm:"type:double precision;not null"`
	MonthRate float64 `json:"month_rate" gorm:"type:double precision;not null"`
	Image     string  `json:"image" gorm:"type:varchar(256);not null"`
	Orders    []Order `json:"orders" gorm:"foreignKey:CarID;constraint:OnDelete:CASCADE"`
}

// CarRequest for handling form data
type CarRequest struct {
	CarName   string                `form:"car_name"`
	DayRate   float64               `form:"day_rate"`
	MonthRate float64               `form:"month_rate"`
	Image     *multipart.FileHeader `form:"image"`
}
