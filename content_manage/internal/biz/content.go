package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Content struct {
	ID             int64         `json:"id"`              // 内容标题
	ContentID      string        `json:"content_id"`      // 内容ID
	Title          string        `json:"title"`           // 内容标题
	VideoURL       string        `json:"video_url"`       // 视频播放URL
	Author         string        `json:"author"`          // 作者
	Description    string        `json:"description"`     // 内容描述
	Thumbnail      string        `json:"thumbnail"`       // 封面图URL
	Category       string        `json:"category"`        // 内容分类
	Duration       time.Duration `json:"duration"`        // 内容时长
	Resolution     string        `json:"resolution"`      // 分辨率 如720p、1080p
	FileSize       int64         `json:"fileSize"`        // 文件大小
	Format         string        `json:"format"`          // 文件格式 如MP4、AVI
	Quality        int32         `json:"quality"`         // 视频质量 1-高清 2-标清
	ApprovalStatus int32         `json:"approval_status"` // 审核状态 1-审核中 2-审核通过 3-审核不通过
	UpdatedAt      time.Time     `json:"updated_at"`      // 内容更新时间
	CreatedAt      time.Time     `json:"created_at"`      // 内容创建时间
}

// 增删改查，查的Find的查找的【参数定义】
type FindParams struct {
	ID       int64
	Author   string
	Title    string
	Page     int32
	PageSize int32
}

type ContentRepo interface {
	Create(context.Context, *Content) error
	Update(context.Context, int64, *Content) error
	IsExist(context.Context, int64) (bool, error)
	Delete(context.Context, int64) error
	Find(context.Context, *FindParams) ([]*Content, int64, error)
}

// ContentUsecase is a Content usecase.
type ContentUsecase struct {
	repo ContentRepo
	log  *log.Helper
}

// NewContentUsecase new a Content usecase.
func NewContentUsecase(repo ContentRepo, logger log.Logger) *ContentUsecase {
	return &ContentUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateContent creates a Content, and returns the new Content.
func (uc *ContentUsecase) CreateContent(ctx context.Context, c *Content) error {
	uc.log.WithContext(ctx).Infof("CreateContent: %v", c)
	return uc.repo.Create(ctx, c)
}

// UpdateContent update a Content.
func (uc *ContentUsecase) UpdateContent(ctx context.Context, c *Content) error {
	uc.log.WithContext(ctx).Infof("UpdateContent: %+v", c)
	return uc.repo.Update(ctx, c.ID, c)
}

// DeleteContent delete a Content.
func (uc *ContentUsecase) DeleteContent(ctx context.Context, id int64) error {
	ok, err := uc.repo.IsExist(ctx, id)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("内容不存在,无法删除")
	}
	//内容存在，执行删除操作
	return uc.repo.Delete(ctx, id)
}

// FindContent find Content.
func (uc *ContentUsecase) FindContent(ctx context.Context, params *FindParams) ([]*Content, int64, error) {
	repo := uc.repo
	contents, total, err := repo.Find(ctx, params) //调用data层的Find实现
	if err != nil {
		return nil, 0, err
	}
	return contents, total, nil
}
