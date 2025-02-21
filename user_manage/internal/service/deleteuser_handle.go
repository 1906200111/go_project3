package service

import (
	"context"
	"user_manage/api/operate"
)

func (a *UserService) DeleteUser(ctx context.Context, req *operate.DeleteUserRequest) (*operate.DeleteUserReply, error) {
	err := a.uc.DeleteUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return nil, nil
}
