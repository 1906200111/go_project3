package service

import (
	"content_manage/api/operate"
	"content_manage/internal/biz"
)

type AppService struct {
	operate.UnimplementedAppServer

	uc *biz.GreeterUsecase
}

// NewAppService
func NewAppService(uc *biz.GreeterUsecase) *AppService {
	return &AppService{uc: uc}
}
