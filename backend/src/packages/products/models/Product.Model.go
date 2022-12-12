package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string `json:"Name"`
	Price int    `json:"Price"`
}
