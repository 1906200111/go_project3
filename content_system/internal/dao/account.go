package dao

import (
	"content_system/internal/model"
	"fmt"
	"gorm.io/gorm"
)

type AccountDao struct {
	db *gorm.DB
}

func NewAccountDao(db *gorm.DB) *AccountDao {
	return &AccountDao{db: db}
}

// 判断这个用户是否存在
func (a *AccountDao) IsExist(userID string) (bool, error) {
	var account model.Account                                      //定义Account实例
	err := a.db.Where("user_id = ?", userID).First(&account).Error //如果找到这个用户，就赋值给这个Account实例
	if err == gorm.ErrRecordNotFound {                             //如果没找到这个用户
		return false, nil
	}
	if err != nil { //如果出现了其他致命错误
		fmt.Printf("IsExist error = [%v] \n", err)
		return false, err
	}
	return true, nil //如果找到了该用户

}

// 创建用户
func (a *AccountDao) Create(account model.Account) error {
	if err := a.db.Create(&account).Error; err != nil {
		fmt.Printf("Create Account error = [%v] \n", err)
		return err
	}
	return nil
}

// 通过id查询这个人的信息
func (a *AccountDao) GetInfoByUserID(userID string) (*model.Account, error) {
	var account model.Account
	err := a.db.Where("user_id = ?", userID).First(&account).Error
	if err != nil {
		fmt.Printf("GetInfoByUserID error = [%v] \n", err)
		return nil, err

	}
	return &account, nil
}
