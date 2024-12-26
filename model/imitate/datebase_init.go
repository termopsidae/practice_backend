package imitate

import (
	"gorm.io/gorm"
	"paractice/model"
	"paractice/pkg"
	"strconv"
	"time"
)

func insertImitateManager(tx *gorm.DB) error {
	manager := model.Manager{
		UserName: "admin",
		Password: "8df52d713468c0ce210209fc3ceacbcf",
		Class:    "1",
		Token:    pkg.RandomString(64) + ":" + strconv.FormatInt(time.Now().Unix(), 10),
		Flag:     "1",
	}
	_, err := manager.InsertNewManager(tx)
	if err != nil {
		return err
	}

	return nil
}
func insertImitateGoods(tx *gorm.DB) error {
	good := model.Good{
		Name:        "test",
		Price:       100,
		Description: "test",
		LastAmount:  100,
		Flag:        "1",
	}
	_, err := good.InsertNewGood(tx)
	if err != nil {
		return err
	}
	return nil
}
