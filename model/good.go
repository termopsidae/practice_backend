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
