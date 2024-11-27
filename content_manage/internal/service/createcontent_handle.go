package service

import (
	"content_manage/api/operate"
	"content_manage/internal/biz"
	"context"
	"github.com/google/uuid"
	"time"
)

func (a *AppService) CreateContent(ctx context.Context,
	req *operate.CreateContentReq) (*operate.CreateContentRsp, error) {
	content := req.GetContent()
	uc := a.uc
	contentID := uuid.New().String()
	err := uc.CreateContent(ctx, &biz.Content{
		ContentID:      contentID,
		Title:          content.GetTitle(),
		VideoURL:       content.GetVideoUrl(),
		Author:         content.GetAuthor(),
		Description:    content.GetDescription(),
		Thumbnail:      content.GetThumbnail(),
		Category:       content.GetCategory(),
		Duration:       time.Duration(content.GetDuration()),
		Resolution:     content.GetResolution(),
		FileSize:       content.GetFileSize(),
		Format:         content.GetFormat(),
		Quality:        content.GetQuality(),
		ApprovalStatus: content.GetApprovalStatus(),
	})
	if err != nil {
		return nil, err
	}
	return &operate.CreateContentRsp{}, nil
}
