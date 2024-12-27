package model

import (
	"gorm.io/gorm"
)

// Cart struct
type Cart struct {
	gorm.Model
	UserId          uint
	User            User
	GoodOrderIdList string
	TotalPrice      float64
	Flag            string
}

// 通过 userId 查询 Card 购物车
func SelectCartByUserId(db *gorm.DB, userId uint) (cart Cart, err error) {
	err = db.Model(&cart).Where("user_id = ?", userId).First(&cart).Error
	return
}

// 为当前用户创建购物车
func (cart *Cart) InsertCart(db *gorm.DB) (id uint, err error) {
	result := db.Create(cart)
	if result.Error != nil {
		return 0, result.Error
	} else {
		return cart.ID, nil
	}
}
func (c *Cart) GetById(db *gorm.DB) error {
	return db.First(&c, c.ID).Error
}

// UpdateCart 更新购物车
func UpdateCart(db *gorm.DB, id uint, params map[string]interface{}) error {
	err := db.Model(&Cart{}).Where("id = ?", id).Updates(params).Error
	if err != nil {
		return err
	}
	return nil
}
