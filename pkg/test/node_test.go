package test

import (
	"paractice/database"
	"paractice/model"
	"paractice/pkg"
	"testing"
)

func TestAddOrder(t *testing.T) {
	database.ConnectDB()
	nodes, err := model.SelectAllNodeOrder(database.DB)
	if err != nil {
		println(err)
	}
	for _, n := range nodes {
		//赠送矿机
		mining := model.Mining{}
		mining.ID = n.Node.MiningId
		err = mining.GetById(database.DB)
		if err != nil {
			println(err)
			return
		}
		newMiningOrder := model.MiningOrder{
			UserId:       n.UserId,
			MiningId:     mining.ID,
			ReleaseNum:   mining.AllIncome,
			ReleaseLimit: mining.AllIncome,
			Flag:         "1",
		}
		_, err = newMiningOrder.InsertNewMiningOrder(database.DB)
		if err != nil {
			println(err)
			return
		}
		user := model.User{}
		user.ID = n.UserId
		err = user.GetById(database.DB)
		if err != nil {
			println(err)
			return
		}
		params := map[string]interface{}{"b_balance": pkg.BigIntStringAdd(user.BBalance, n.Node.BBenefit)}
		err = model.UpdateUser(database.DB, user.ID, params)
		if err != nil {
			println(err)
			return
		}
	}
}
