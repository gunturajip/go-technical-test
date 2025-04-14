package Model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Order struct {
	Base
	Address           string `json:"address" form:"address" valid:"-"`
	PurchaseProofLink string `json:"purchase_proof_link" form:"purchase_proof_link" valid:"-"`
	Status            string `gorm:"not null" json:"status"`
	UserID            uint   `gorm:"not null" json:"user_id"`
	Items             []Item `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"items"`
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
