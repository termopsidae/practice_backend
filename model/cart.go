package model

import (
	"gorm.io/gorm"
)

// Cart struct
type Cart struct {
	gorm.Model
	UserId          uint
	User            User
	GoodOrderIdList []uint
	TotalPrice      float64
	Flag            string
}

func (c *Cart) GetById(db *gorm.DB) error {
	return db.First(&c, c.ID).Error
}
