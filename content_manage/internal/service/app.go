package service

import (
	"content_manage/api/operate"
	"content_manage/internal/biz"
)

type AppService struct {
	operate.UnimplementedAppServer

	uc *biz.ContentUsecase
}

// NewAppService
func NewAppService(uc *biz.ContentUsecase) *AppService {
	return &AppService{uc: uc}
}
