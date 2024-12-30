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

// 插入订单数据
// InsertNewUser 新增订单
func (goodOrder *GoodOrder) InsertNewGoodOrder(db *gorm.DB) (id uint, err error) {
	result := db.Create(goodOrder)
	if result.Error != nil {
		return 0, result.Error
	} else {
		return goodOrder.ID, nil
	}
}

// 查询订单，通过IDs ，和购买状态
func SelectOrderByIdsAndFlag(db *gorm.DB, ids []uint, flag string) (goodOrders []GoodOrder, err error) {
	tx := db.Model(&GoodOrder{})
	if flag != "" { //查询当前IDs ，未付款订单
		err = tx.Where("id IN (?)", ids).Where("flag = ?", flag).Find(&goodOrders).Error
	} else { //查询当前IDs所有订单
		err = tx.Where("id IN (?)", ids).Find(&goodOrders).Error
	}
	return goodOrders, err
}

// 查询订单 通过Id,和购买状态
func SelectOrderByIdAndFlag(db *gorm.DB, id uint, flag string) (goodOrder GoodOrder, err error) {
	tx := db.Model(&GoodOrder{})
	if flag != "" { //查询当前IDs ，未付款订单
		err = tx.Where("id = ?", id).Where("flag = ?", flag).First(&goodOrder).Error
	} else { //查询当前IDs所有订单
		err = tx.Where("id = ?", id).First(&goodOrder).Error
	}
	return goodOrder, err
}

// UpdateCart 更新订单数据
func UpdateGoodOrder(db *gorm.DB, id uint, params map[string]interface{}) error {
	err := db.Model(&GoodOrder{}).Where("id = ?", id).Updates(params).Error
	if err != nil {
		return err
	}
	return nil
}
