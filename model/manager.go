package model

import (
	"errors"
	"gorm.io/gorm"
)

// Manager struct
type Manager struct {
	gorm.Model
	UserName string
	Password string
	Class    string // 等级(1-主要 2-普通)
	Token    string
	Flag     string // 启用标志(1-启用 0-停用)
}

func (p *Manager) GetById(db *gorm.DB) error {
	return db.First(&p, p.ID).Error
}

func QueryManagerCount(db *gorm.DB) (uCount int64, err error) {
	if err := db.Model(&Manager{}).Count(&uCount).Error; err != nil {
		return 0, err
	}
	return uCount, nil
}
func (m *Manager) GetByUserName(db *gorm.DB) error {
	return db.Model(&m).Where("user_name = ? ", m.UserName).Take(&m).Error
}

// SelectAllManager 查询所有管理员
func SelectAllManager(db *gorm.DB, class string) (ps []Manager, err error) {
	if class == "1" {
		if err := db.Model(&Manager{}).Order("id").Find(&ps).Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.Model(&Manager{}).Where("class = ?", class).Order("id").Find(&ps).Error; err != nil {
			return nil, err
		}
	}
	return ps, nil
}

// SelectManagerById 查询管理员通过ID
func SelectManagerById(db *gorm.DB, managerId uint) (manager Manager, err error) {
	err = db.Model(&manager).Where("id = ?", managerId).First(&manager).Error
	return manager, err
}

// SelectAllManagerID SelectAllManager 查询所有管理员ID
func SelectAllManagerID(db *gorm.DB) (ps []uint, err error) {
	ps = make([]uint, 0)
	if err := db.Model(&Manager{}).Select("id").Order("id").Find(&ps).Error; err != nil {
		return nil, err
	}
	return ps, nil
}

// InsertNewManager 新增管理员
func (m *Manager) InsertNewManager(db *gorm.DB) (id uint, err error) {
	result := db.Create(m)
	if result.Error != nil {
		return 0, result.Error
	} else {
		return m.ID, nil
	}
}

// UpdateManagerFlag 更新管理员Flag
func (m *Manager) UpdateManagerFlag(db *gorm.DB) error {
	res := db.Model(&m).Where("id = ?", m.ID).Update("flag", m.Flag)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected == 0")
	}
	return nil
}

// UpdateManagerToken 更新用户Token
func (m *Manager) UpdateManagerToken(db *gorm.DB) error {
	res := db.Model(&m).Where("id = ?", m.ID).Update("token", m.Token)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected == 0")
	}
	return nil
}

// ManagerSelectIdByToken token查询用户数据 token = "HASH"
func ManagerSelectIdByToken(db *gorm.DB, token string) (managerId uint, tokenData string, err error) {
	err = db.Model(&Manager{}).
		Select("id", "token").
		Where("token LIKE ?", token+":%").
		Row().Scan(&managerId, &tokenData)
	return
}

// ManagerSelectByToken token查询用户数据 token = "HASH"
func ManagerSelectByToken(db *gorm.DB, token string) (manager Manager, err error) {
	err = db.Where("token LIKE ?", token+":%").First(&manager).Error
	return
}

// ManagerRefreshToken
// @Description: 修改指定用户的token数据
// @param token 数据格式 <token_value:timestamp>
// @return err
func ManagerRefreshToken(db *gorm.DB, managerId uint, token string) (err error) {
	res := db.Model(&Manager{}).Where("id = ?", managerId).Update("token", token)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected == 0")
	}
	return nil
}

// UpdateManager 更新管理员
func UpdateManager(db *gorm.DB, id uint, params map[string]interface{}) error {
	err := db.Model(&Manager{}).Where("id = ?", id).Updates(params).Error
	if err != nil {
		return err
	}
	return nil
}
