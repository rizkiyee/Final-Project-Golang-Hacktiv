package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	User    *User  `json:"user,omitempty"`
	Photo   *Photo `json:"photo,omitempty"`
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
	Message string 	`gorm:"not null" form:"message" json:"message" valid:"required~Message is required"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
