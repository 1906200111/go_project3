package service

import (
	"context"
	"user_manage/api/operate"
	"user_manage/internal/biz"
)

func (a *UserService) CreateUser(ctx context.Context, req *operate.CreateUserRequest) (*operate.CreateUserReply, error) {
	user := req.GetUser()
	err := a.uc.CreateUser(ctx, &biz.User{
		UserID:   user.GetUserid(),
		Password: user.GetPassword(),
		Nickname: user.GetNickname(),
	})
	if err != nil {
		return nil, err
	}
	return &operate.CreateUserReply{}, nil
}
