package Model

import (
	"1/Util"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	Base
	Username string  `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email    string  `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password string  `gorm:"not null" json:"password,omitempty" form:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Admin    bool    `gorm:"not null;default:false" json:"admin" form:"admin" valid:"-"`
	Orders   []Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return err
	}

	u.Password = Util.HashPass(u.Password)
	err = nil
	return err
}
