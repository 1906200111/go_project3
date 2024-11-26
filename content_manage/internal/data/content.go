package data

import (
	"content_manage/internal/biz"
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

// NewGreeterRepo .
func NewContentRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
