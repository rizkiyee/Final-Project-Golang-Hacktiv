package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string    `gorm:"not null" form:"title" json:"title" valid:"required~Title is required"`
	Caption  string    `form:"caption" json:"caption"`
	PhotoUrl string    `gorm:"not null" form:"photo_url" json:"photo_url" valid:"required~Photo URL is required"`
	UserID   uint      `json:"user_id"`
	Users    *User     `gorm:"ForeignKey:UserID" json:"user,omitempty"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments,omitempty"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
