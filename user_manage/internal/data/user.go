package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"user_manage/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type UserDetail struct {
	gorm.Model
	UserID   string `gorm:"column:user_id;"`
	Password string `gorm:"column:password;"`
	Nickname string `gorm:"column:nickname;"`
}

func (UserDetail) TableName() string {
	table := "cms_content.user"
	return table
}

func (c *userRepo) Create(ctx context.Context, user *biz.User) error {
	c.log.Infof("userRepo Create user = %+v", user)
	detail := UserDetail{
		UserID:   user.UserID,
		Password: user.Password,
		Nickname: user.Nickname,
	}
	db := c.data.db
	if err := db.Create(&detail).Error; err != nil {
		c.log.Errorf("user create error = %v", err)
		return err
	}

	return nil
}

func (c *userRepo) Update(ctx context.Context, id int64, user *biz.User) error {
	c.log.Infof("userRepo Update user = %+v", user)
	detail := UserDetail{
		UserID:   user.UserID,
		Password: user.Password,
		Nickname: user.Nickname,
	}
	db := c.data.db
	if err := db.Where("id = ?", id).Updates(&detail).Error; err != nil {
		c.log.WithContext(ctx).Errorf("user update error = %v", err)
		return err
	}
	return nil
}

func (c *userRepo) IsExist(ctx context.Context, id int64) (bool, error) {
	db := c.data.db
	var detail UserDetail
	err := db.Where("id = ?", id).First(&detail).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		c.log.WithContext(ctx).Errorf("User isExist = [%v]", err)
		return false, err
	}
	return true, nil
}

func (c *userRepo) Delete(ctx context.Context, id int64) error {
	db := c.data.db
	// 删除索引信息
	err := db.Where("id = ?", id).
		Delete(&UserDetail{}).Error
	if err != nil {
		c.log.WithContext(ctx).Errorf("user delete error = %v", err)
		return err
	}
	return nil
}

func (c *userRepo) Find(ctx context.Context, params *biz.FindParams) ([]*biz.User, int64, error) {
	query := c.data.db.Model(&UserDetail{})
	// 构造查询条件
	if params.ID != 0 {
		query = query.Where("id = ?", params.ID)
	}
	if params.UserID != "" {
		query = query.Where("user_id = ?", params.UserID)
	}
	if params.Nickname != "" {
		query = query.Where("nickname = ?", params.Nickname)
	}
	// 总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	//设置默认页大小
	var page, pageSize = 1, 10
	if params.Page > 0 {
		page = int(params.Page)
	}
	if params.Page_Size > 0 {
		pageSize = int(params.Page_Size)
	}
	offset := (page - 1) * pageSize
	var results []*UserDetail
	if err := query.Offset(offset).Limit(pageSize).Find(&results).Error; err != nil {
		c.log.WithContext(ctx).Errorf("user find error = %v", err)
		return nil, 0, err
	}
	var users []*biz.User
	for _, r := range results {
		users = append(users, &biz.User{
			ID:       int64(r.ID),
			UserID:   r.UserID,
			Nickname: r.Nickname,
		})
	}
	return users, total, nil
}

//操作db
