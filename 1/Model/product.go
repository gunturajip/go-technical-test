package Model

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	Base
	Name  string `gorm:"not null" json:"name" form:"address" valid:"required~Name of your product is required"`
	Price uint   `gorm:"not null" json:"price" form:"price" valid:"required~Price of your product is required, numeric~Price must be numeric"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return err
	}

	if p.Price <= 0 {
		return errors.New("price must be over 0")
	}

	err = nil
	return err
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return err
	}

	if p.Price <= 0 {
		return errors.New("price must be over 0")
	}

	err = nil
	return err
}
