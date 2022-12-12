package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     *string
	UserType  string
	Password  string
}
