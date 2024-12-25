package api

import (
	"gorm.io/gorm"
	"paractice/model"
	"time"
)

type SingleUser struct {
	Id            uint
	RecommenderId uint
	Branch        []uint
}

var UserTree map[uint]SingleUser
var LockedMap map[uint]LockedQueue
var LockedAMap map[uint]LockedQueue
var LockedBMap map[uint]LockedQueue

type LockedQueue struct {
	AllAmount    string
	SingleLocked map[time.Time]string
}

// AddNewBranch 增加新的用户树枝干 treeId:推荐人 branchId:当前用户
func AddNewBranch(treeId uint, branchId uint) {
	branch := make([]uint, 0)
	UserTree[branchId] = SingleUser{
		Id:            branchId,
		RecommenderId: treeId,

		Branch: branch,
	}
	//newBranch := make([]uint, 0)
	newBranch := append(UserTree[treeId].Branch, branchId)
	UserTree[treeId] = SingleUser{
		Id:            treeId,
		RecommenderId: UserTree[treeId].RecommenderId,

		Branch: newBranch,
	}
}

// IncreaseTeamNumber 增加用户的团队用户人数
func IncreaseTeamNumber(Id uint) {
	UserTree[Id] = SingleUser{
		Id:            Id,
		RecommenderId: UserTree[Id].RecommenderId,

		Branch: UserTree[Id].Branch,
	}
}

// GetAssociate 用户下级激活用户
func GetAssociate(walletId uint, tx *gorm.DB) int64 {
	var count int64
	for _, b := range UserTree[walletId].Branch {
		u := model.User{}
		u.ID = b
		err := u.GetById(tx)
		if err != nil {
			return 0
		}
		if u.Flag == "1" {
			count += 1
		}
	}
	return count
}

// InitUserTree 初始化用户树
func InitUserTree(db *gorm.DB) error {
	UserTree = make(map[uint]SingleUser)
	LockedMap = make(map[uint]LockedQueue)
	LockedAMap = make(map[uint]LockedQueue)
	LockedBMap = make(map[uint]LockedQueue)
	// 查询出所有用户
	users, err := model.SelectAllUser(db)
	if err != nil {
		return err
	}
	for _, user := range users {
		if user.ID == 0 {
			continue
		}
		// 如果用户推荐人是自己，就修改为0
		if user.RecommendId == user.ID {
			params := map[string]interface{}{"recommend_id": 0}
			err = model.UpdateUser(db, user.ID, params)
			if err != nil {
				return err
			}
		}
		lq := make(map[time.Time]string)
		LockedMap[user.ID] = LockedQueue{AllAmount: "0", SingleLocked: lq}
		LockedAMap[user.ID] = LockedQueue{AllAmount: "0", SingleLocked: lq}
		LockedBMap[user.ID] = LockedQueue{AllAmount: "0", SingleLocked: lq}
		UserTree[user.ID] = SingleUser{
			Id:            user.ID,
			RecommenderId: user.RecommendId,

			Branch: nil,
		}

	}

	initUserBranch(users)
	return nil
}

// initUserBranch 内部方法 用于初始化用户数的下级用户数列添加
func initUserBranch(users []model.User) {
	for _, u := range users {
		if u.ID == 0 {
			continue
		}
		newBranch := make([]uint, 0)
		if UserTree[u.RecommendId].Branch != nil {
			newBranch = append(UserTree[u.RecommendId].Branch, u.ID)
		} else {
			newBranch = append(newBranch, u.ID)
		}

		UserTree[u.RecommendId] = SingleUser{
			Id:            UserTree[u.ID].Id,
			RecommenderId: UserTree[u.RecommendId].RecommenderId,

			Branch: newBranch,
		}
	}
}
