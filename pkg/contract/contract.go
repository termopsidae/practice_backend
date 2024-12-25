package contract

import (
	"fmt"
	"gorm.io/gorm"
	"paractice/model"
	"paractice/model/api"
	"paractice/pkg"
)

//func CheckTransaction() {
//	for {
//		err := database.DB.Transaction(func(tx *gorm.DB) error {
//			cs, err := model.QueryAllTransactionUnChecked(tx)
//			if err != nil {
//				return err
//			}
//			if len(cs) > 0 {
//				for _, txs := range cs {
//					ctx := context.Background()
//					client, err := ethclient.Dial("https://bsc-dataseed.binance.org/") // 更换为你的BSC节点地址
//					if err != nil {
//						return err
//					}
//
//					txHash := common.HexToHash(txs.Hash) // 替换为你的交易哈希
//					_, isPending, err := client.TransactionByHash(ctx, txHash)
//					if err != nil {
//						return err
//					}
//					if isPending {
//
//					} else {
//						params := map[string]interface{}{"flag": "1"}
//						err = model.UpdateTransaction(tx, txs.ID, params)
//						if err != nil {
//							return err
//						}
//						err = model.UpdateMiningOrder(tx, txs.AssociatedId, params)
//						if err != nil {
//							return err
//						}
//						mo := model.MiningOrder{}
//						mo.ID = txs.AssociatedId
//						err = mo.GetById(tx)
//						if err != nil {
//							return err
//						}
//						consumer := model.User{}
//						consumer.ID = mo.UserId
//						err = consumer.GetById(tx)
//						if err != nil {
//							return err
//						}
//						if consumer.RecommendId != 0 {
//							directRecommender := model.User{}
//							directRecommender.ID = consumer.RecommendId
//							err = directRecommender.GetById(tx)
//							if err != nil {
//								return err
//							}
//							err := UpdateTeamPerformance(tx, consumer.ID, directRecommender, mo.Mining.Price)
//							if err != nil {
//								return err
//							}
//						}
//					}
//				}
//			}
//
//			return nil
//		})
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//		time.Sleep(30 * time.Second)
//
//	}
//}

// UpdateTeamPerformance 更新团队业绩 id:用户id directlyRecommender:推荐人 performance:本次消费
func UpdateTeamPerformance(tx *gorm.DB, directlyRecommender model.User) (err error) {
	var nowPerformance = "0"
	finalLevel := directlyRecommender.Level
	var performanceMap map[int64]string
	if len(performanceMap) == 0 {
		performanceMap, err = getPerformanceMap(tx)
		if err != nil {
			return err
		}
	}
	maxBranchPerformance := "0"
	branches := api.UserTree[directlyRecommender.ID].Branch
	for _, branch := range branches {
		b := model.User{}
		b.ID = branch
		err = b.GetById(tx)
		if err != nil {
			return err
		}
		inPerformance := pkg.BigIntStringAdd(b.Performance, b.Consume)
		tempPerformance := "0"
		tempPerformance = pkg.BigIntStringAdd(nowPerformance, inPerformance)
		nowPerformance = tempPerformance
		if pkg.CmpBigIntString(inPerformance, maxBranchPerformance) == 1 {
			maxBranchPerformance = inPerformance
		}
	}
	params := map[string]interface{}{"performance": nowPerformance}
	params["small_zone_performance"] = pkg.BigIntStringSub(nowPerformance, maxBranchPerformance)
	for i := directlyRecommender.Level + 1; i < int64(len(performanceMap)); i++ {
		if pkg.CmpBigIntString(pkg.BigIntStringSub(nowPerformance, maxBranchPerformance), performanceMap[i]) != -1 {
			finalLevel = i
			continue
		} else {
			break
		}
	}
	if directlyRecommender.Level < finalLevel {
		params["level"] = finalLevel
	}
	err = model.UpdateUser(tx, directlyRecommender.ID, params)
	if err != nil {
		return err
	}
	//api.IncreaseTeamNumber(directlyRecommender.RecommendAddress)
	if directlyRecommender.RecommendId != 0 {
		superiorRecommender := model.User{}
		superiorRecommender.ID = directlyRecommender.RecommendId
		err = superiorRecommender.GetById(tx)
		if err != nil {
			return err
		}
		err = UpdateTeamPerformance(tx, superiorRecommender)
		if err != nil {
			return err
		}
	}
	return nil
}

// 获取业绩map
func getPerformanceMap(tx *gorm.DB) (outPerformanceMap map[int64]string, err error) {
	var kVMap map[string]string
	// 确保必须有key value map数据
	if len(kVMap) == 0 {
		kVMap, err = getKVMap(tx)
		if err != nil {
			return nil, err
		}
	}
	outPerformanceMap = make(map[int64]string)
	// 循环处理多个键
	for i := 0; i < 8; i++ {
		key := fmt.Sprintf("V%dPerformance", i)
		valStr := kVMap[key]
		if valStr != "" {
			outPerformanceMap[int64(i)] = valStr
		} else {
			outPerformanceMap[int64(i)] = "0"
		}
	}
	return outPerformanceMap, nil
}

// 获取key-value表数据存入map
func getKVMap(tx *gorm.DB) (map[string]string, error) {
	outKVMap, err := model.SelectAllKeyValueToMap(tx)
	if err != nil {
		return nil, fmt.Errorf("getKVMap error1 : %s", err.Error())
	}
	if len(outKVMap) == 0 {
		return nil, fmt.Errorf("getKVMap error2 : have no key value in table")
	}
	return outKVMap, nil
}
