package model

import (
	"gorm.io/gorm"
)

// Good struct
type Good struct {
	gorm.Model
	Name        string
	Price       float64
	Description string
	LastAmount  int64
	Flag        string // 启用标志(0-停用 1-可购买)
}

func (g *Good) GetById(db *gorm.DB) error {
	return db.First(&g, g.ID).Error
}

// InsertNewGood 新增商品
func (g *Good) InsertNewGood(db *gorm.DB) (id uint, err error) {
	result := db.Create(g)
	if result.Error != nil {
		return 0, result.Error

	} else {
		return g.ID, nil
	}
}

// SelectAllGoods 查询所有
func SelectAllGoods(db *gorm.DB, flag string) (us []Good, err error) {
	tx := db.Model(&Good{})
	if flag != "1" {
		if err := tx.Where("flag = ?", "0").Order("created_at desc").Find(&us).Error; err != nil {
			return nil, err
		}
	} else {
		if err := tx.Where("flag = ?", flag).Order("created_at desc").Find(&us).Error; err != nil {
			return nil, err
		}
	}
	return us, nil
}

//查询商品通过 goodId

func SelectGoodById(db *gorm.DB, id uint) (good Good, err error) {
	err = db.Model(&good).Where("id = ?", id).First(&good).Error
	return
}
