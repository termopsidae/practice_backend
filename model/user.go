package model

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"paractice/pkg"
)

// User struct
type User struct {
	gorm.Model
	RecommendId   uint
	WalletAddress string
	Balance       float64
	Token         string
	Flag          string // 启用标志(0-正常 -1-冻结)
}

func (u *User) GetById(db *gorm.DB) error {
	// 获取用户信息
	err := db.Where("id = ?", u.ID).Take(&u).Error
	if err != nil {
		return err
	}

	// 检查 RecommendId 是否和自己的 ID 相同
	if u.RecommendId == u.ID {
		// 将 RecommendId 设置为 0
		u.RecommendId = 0

		// 更新表中的 recommend_id 字段
		err = db.Model(&u).Where("id = ?", u.ID).Update("recommend_id", u.RecommendId).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *User) GetByWalletAddress(db *gorm.DB) error {
	return db.Model(&u).Where("wallet_address = ?", u.WalletAddress).First(&u).Error
}

func GetByWalletAddressLast8(db *gorm.DB, walletAddressLast8 string) (u User, err error) {
	err = db.Model(&u).Where("wallet_address LIKE ?", "%"+walletAddressLast8).First(&u).Error
	return
}

// SelectAllUser 查询所有用户
func SelectAllUser(db *gorm.DB) (us []User, err error) {
	if err := db.Model(&User{}).Order("created_at desc").Find(&us).Error; err != nil {
		return nil, err
	}
	return us, nil
}

// SelectUserTeam 查询查询用户的团队
func SelectUserTeam(db *gorm.DB, recommendId uint) (us []User, err error) {
	us = make([]User, 0)
	if err := db.Model(&User{}).Where("recommend_id = ?", recommendId).Order("created_at desc").Find(&us).Error; err != nil {
		return nil, err
	}
	return us, nil
}

// InsertNewUser 新增用户
func (u *User) InsertNewUser(db *gorm.DB) (id uint, err error) {
	result := db.Create(u)
	if result.Error != nil {
		return 0, result.Error
	} else {
		return u.ID, nil
	}
}

// UpdateUserToken 更新用户Token
func (u *User) UpdateUserToken(db *gorm.DB) error {
	res := db.Model(&u).Where("id = ?", u.ID).Update("token", u.Token)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected == 0")
	}
	return nil
}

func (u *User) UpdateUseRecommender(db *gorm.DB) error {
	res := db.Model(&u).Where("id = ?", u.ID).Update("recommend_id", u.RecommendId)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected == 0")
	}
	return nil
}

// SelectUserByParams 查询用户
func SelectUserByParams(db *gorm.DB, params map[string]interface{}) (us []User, err error) {
	tx := db.Model(&User{})
	if params["flag"] != nil {
		tx.Where("flag = ?", params["flag"])
	}
	if params["recommend_id"] != nil {
		tx.Where("recommend_id = ?", params["recommend_id"])
	}
	if params["wallet_address"] != nil {
		tx.Where("wallet_address = ?", params["wallet_address"])
	}
	if params["order_by"] != nil {
		tx.Order("CAST(" + params["order_by"].(string) + " AS DOUBLE PRECISION) desc")
	} else {
		tx.Order("created_at desc")
	}
	if params["page"] != nil {
		tx.Limit(10).Offset(params["page"].(int) * 10)
	}
	err = tx.Find(&us).Error
	if err != nil {
		return nil, err
	}
	return us, nil
}

// UserSelectAddressByToken token查询用户数据 token = "HASH"
func UserSelectAddressByToken(db *gorm.DB, token string) (id uint, tokenData string, err error) {
	err = db.Model(&User{}).
		Select("id", "token").
		Where("token LIKE ?", token+":%").
		Row().Scan(&id, &tokenData)
	return
}

// UserRefreshToken
// @Description: 修改指定用户的token数据
// @param token 数据格式 <token_value:timestamp>
// @return err
func UserRefreshToken(db *gorm.DB, id uint, token string) (err error) {
	res := db.Model(&User{}).Where("id = ?", id).Update("token", token)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected == 0")
	}
	return nil
}

// UpdateUser 更新用户
func UpdateUser(db *gorm.DB, id uint, params map[string]interface{}) error {
	err := db.Model(&User{}).Where("id = ?", id).Updates(params).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateUserByAddress 更新用户
func UpdateUserByAddress(db *gorm.DB, walletAddress string, params map[string]interface{}) error {
	err := db.Model(&User{}).Where("wallet_address = ?", walletAddress).Updates(params).Error
	if err != nil {
		return err
	}
	return nil
}

// QueryUserCountByParams 条件用户计数
func QueryUserCountByParams(db *gorm.DB, params map[string]interface{}) (uCount int64, err error) {
	tx := db.Model(&User{})
	if params["flag"] != nil {
		tx.Where("flag = ?", params["flag"])
	}
	if params["recommend_id"] != nil {
		tx.Where("recommend_id = ?", params["recommend_id"])
	}
	if params["wallet_address"] != nil {
		tx.Where("wallet_address = ?", params["wallet_address"])
	}
	if params["level"] != nil {
		tx.Where("level = ?", params["level"])
	}
	if err := tx.Count(&uCount).Error; err != nil {
		return 0, err
	}
	return uCount, nil
}

/*
UpdateUserBalance
修改用户的 余额
agoUsdtBalance = 原来db的数据
newUsdtBalance = 新的db数据
*/
func UpdateUserBalance(db *gorm.DB, userID uint, agoUsdtBalance, newUsdtBalance string) error {
	if pkg.CmpBigIntString(newUsdtBalance, "0") == -1 {
		return errors.New("update failed: balance can not <0")
	}
	// 加锁获取当前 usdt_balance
	var user User
	if err := db.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ? AND balance = ?", userID, agoUsdtBalance).
		First(&user).Error; err != nil {
		return err
	}

	// 更新余额
	res := db.Model(&user).Update("balance", newUsdtBalance)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("update failed: balance may have been modified by another transaction")
	}

	return nil
}
