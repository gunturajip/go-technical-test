package Model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Order struct {
	Base
	Address       string `gorm:"not null" json:"address" form:"address" valid:"required~Address of your order is required"`
	PaymentMethod string `gorm:"not null" json:"payment_method" form:"payment_method" valid:"required~Payment method of your order is required"`
	UserID        uint   `gorm:"not null" json:"user_id"`
	Items         []Item `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(o)

	if errCreate != nil {
		err = errCreate
		return err
	}

	err = nil
	return err
}

func (o *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(o)

	if errUpdate != nil {
		err = errUpdate
		return err
	}

	err = nil
	return err
}
