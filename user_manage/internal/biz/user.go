package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type User struct {
	ID       int64     `json:"id"`
	UserID   string    `json:"user_id"`
	Password string    `json:"password"`
	Nickname string    `json:"nickname"`
	Ct       time.Time `json:"created_at"` // 内容更新时间
	Ut       time.Time `json:"updated_at"` // 内容创建时间
}

type UserRepo interface {
	Create(context.Context, *User) error
	Update(context.Context, int64, *User) error
	IsExist(context.Context, int64) (bool, error)
	Delete(context.Context, int64) error
	Find(context.Context, *FindParams) ([]*User, int64, error)
}

// 查找的参数
type FindParams struct {
	ID        int64
	UserID    string
	Nickname  string
	Page      int32
	Page_Size int32
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, g *User) error {
	uc.log.WithContext(ctx).Infof("CreateUser: %+v", g)
	return uc.repo.Create(ctx, g)
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, g *User) error {
	uc.log.WithContext(ctx).Infof("UpdateUser: %+v", g)
	return uc.repo.Update(ctx, g.ID, g)
}

func (uc *UserUsecase) DeleteUser(ctx context.Context, id int64) error {
	uc.log.WithContext(ctx).Infof("DeleteUser: %+v", id)
	//复合操作，现判断该用户是否存在，再执行删除操作
	ok, err := uc.repo.IsExist(ctx, id)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("User no exist")
	}
	//用户存在的情况,执行删除操作
	return uc.repo.Delete(ctx, id)
}

func (uc *UserUsecase) FindUser(ctx context.Context, params *FindParams) ([]*User, int64, error) {
	users, total, err := uc.repo.Find(ctx, params)
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

//执行组合逻辑
