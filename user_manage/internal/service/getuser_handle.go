package service

import (
	"context"
	"user_manage/api/operate"
	"user_manage/internal/biz"
)

func (a *UserService) GetUser(ctx context.Context, req *operate.GetUserRequest) (*operate.GetUserReply, error) {
	users, total, err := a.uc.FindUser(ctx, &biz.FindParams{
		Page:      req.GetPage(),
		Page_Size: req.GetPageSize(),
		UserID:    req.GetUserid(),
		Nickname:  req.GetNickname(),
		ID:        req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	var usersInfo []*operate.UserInfo
	for _, user := range users {
		usersInfo = append(usersInfo, &operate.UserInfo{
			Id:       user.ID,
			Userid:   user.UserID,
			Nickname: user.Nickname,
		})
	}
	return &operate.GetUserReply{
		Total: total,
		User:  usersInfo,
	}, nil

}
