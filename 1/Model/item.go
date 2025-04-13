package Model

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Item struct {
	Base
	OrderID   uint    `gorm:"not null" json:"order_id" form:"order_id"`
	ProductID uint    `gorm:"not null" json:"product_id" form:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  uint    `gorm:"not null" json:"quantity" form:"quantity" valid:"required~Quantity of your item is required, numeric~Quantity must be numeric"`
}

func (i *Item) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(i)

	if errCreate != nil {
		err = errCreate
		return err
	}

	if i.Quantity <= 0 {
		return errors.New("quantity must be over 0")
	}

	err = nil
	return err
}

func (i *Item) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(i)

	if errUpdate != nil {
		err = errUpdate
		return err
	}

	if i.Quantity <= 0 {
		return errors.New("quantity must be over 0")
	}

	err = nil
	return err
}
