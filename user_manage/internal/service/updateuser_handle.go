package service

import (
	"context"
	"user_manage/api/operate"
	"user_manage/internal/biz"
)

func (a *UserService) UpdateUser(ctx context.Context, req *operate.UpdateUserRequest) (*operate.UpdateUserReply, error) {
	user := req.GetUser()
	err := a.uc.UpdateUser(ctx, &biz.User{
		ID:       user.GetId(),
		UserID:   user.GetUserid(),
		Password: user.GetPassword(),
		Nickname: user.GetNickname(),
	})
	if err != nil {
		return nil, err
	}
	return &operate.UpdateUserReply{}, nil
}
