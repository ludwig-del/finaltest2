package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string `valid:"required~Name is required"`
	Price float64 `valid:"required~Price must be greater than 0,range(0|999999)~Price must be greater than 0"`
	Quantity int `valid:"required~Quantity is required,range(0|99999)~Quantity cannot be negative"`
	Description string `valid:"required~Description is required,stringlength(10|9999)~Description must be at least 10 characters"`
}

func (s *Product) Validate() error {
	_, err := govalidator.ValidateStruct(s)
	return err
}