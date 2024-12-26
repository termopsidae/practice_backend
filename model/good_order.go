package model

import (
	"gorm.io/gorm"
)

// GoodOrder struct
type GoodOrder struct {
	gorm.Model
	UserId     uint
	GoodId     uint
	Good       Good
	CartId     uint
	Amount     int64
	TotalPrice float64
	Flag       string // 启用标志(0-取消 1-待付款 2-已付款)
}

func (g *GoodOrder) GetById(db *gorm.DB) error {
	return db.First(&g, g.ID).Error
}

// SelectAllGoodOrder 根据状态查询用户所有订单
func SelectAllGoodOrder(db *gorm.DB, flag string) (pos []GoodOrder, err error) {
	if err = db.Model(&GoodOrder{}).Preload("Good").Where("flag IN (?)", flag).Order("id").Find(&pos).Error; err != nil {
		return nil, err
	}
	return pos, nil
}
