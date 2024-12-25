package api

import (
	"fmt"
	"strings"

	//"fmt"
	"paractice/database"
	"paractice/model"
	"paractice/pkg"

	//"paractice/model/imitate"
	"testing"
)

func TestRunP(t *testing.T) {
	database.ConnectDB()
	user := model.User{WalletAddress: pkg.Upper("0XF7AF71D187599517EECDB31B5F98EB8E9D9E54BD")}
	err := user.GetByWalletAddress(database.DB)
	fmt.Println(err)
	fmt.Println(user)
	//IncomeRunP(database.DB)
	//err := imitate.InsertImitateKeyValue(database.DB)
	//fmt.Println(err)
	//if err != nil {
	//	return err
	//}
	//IncomeRunP(database.DB)
	//ContractCycle(database.DB)
}
func TestCut(t *testing.T) {
	Str := "-1100"
	alice := strings.Split(Str, "-")
	println(alice[1])
}
