package models

import (
	"time"

	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	Price int
	Date  time.Time
	// Customer User `gorm:embedded`
	// Product  User `gorm:embedded`
}
