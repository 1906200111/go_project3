package service

import (
	"user_manage/api/operate"
	"user_manage/internal/biz"
)

type UserService struct {
	operate.UnimplementedUserServer
	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}
