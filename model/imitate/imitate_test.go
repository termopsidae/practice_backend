package imitate

import (
	"fmt"
	"paractice/database"
	"paractice/model"
	"testing"
)

func TestDataBase(t *testing.T) {
	InitDataBase()
	_, err := model.SelectAllUser(database.DB)
	if err != nil {
		fmt.Println("----------------------------------------------------")
		fmt.Println(err)
		fmt.Println("----------------------------------------------------")
		return
	}
	fmt.Println("----------------------------------------------------")
}
func TestInsert(t *testing.T) {
	database.ConnectDB()
	err := insertImitateGoods(database.DB)
	if err != nil {
		fmt.Println(err)
	}
	err = insertImitateManager(database.DB)
	if err != nil {
		fmt.Println(err)
	}
}
