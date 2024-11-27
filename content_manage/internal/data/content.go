package data

import (
	"content_manage/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

type contentRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewContentRepo(data *Data, logger log.Logger) biz.ContentRepo {
	return &contentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type ContentDetail struct {
	ID             int64         `gorm:"column:id;primary_key"`  // 自增ID
	ContentID      string        `gorm:"column:content_id"`      // 内容ID
	Title          string        `gorm:"column:title"`           // 内容标题
	Description    string        `gorm:"column:description"`     // 内容描述
	Author         string        `gorm:"column:author"`          // 作者
	VideoURL       string        `gorm:"column:video_url"`       // 视频播放URL
	Thumbnail      string        `gorm:"column:thumbnail"`       // 封面图URL
	Category       string        `gorm:"column:category"`        // 内容分类
	Duration       time.Duration `gorm:"column:duration"`        // 内容时长
	Resolution     string        `gorm:"column:resolution"`      // 分辨率 如720p、1080p
	FileSize       int64         `gorm:"column:fileSize"`        // 文件大小
	Format         string        `gorm:"column:format"`          // 文件格式 如MP4、AVI
	Quality        int32         `gorm:"column:quality"`         // 视频质量 1-高清 2-标清
	ApprovalStatus int32         `gorm:"column:approval_status"` // 审核状态 1-审核中 2-审核通过 3-审核不通过
	UpdatedAt      time.Time     `gorm:"column:updated_at"`      // 内容更新时间
	CreatedAt      time.Time     `gorm:"column:created_at"`      // 内容创建时间
}

func (c ContentDetail) TableName() string {
	return "cms_content.t_content_details"
}

func (c *contentRepo) Create(ctx context.Context, content *biz.Content) error {
	c.log.Infof("contentRepo Create context = %+v", content)
	detail := ContentDetail{
		Title:          content.Title,
		ContentID:      content.ContentID,
		Description:    content.Description,
		Author:         content.Author,
		VideoURL:       content.VideoURL,
		Thumbnail:      content.Thumbnail,
		Category:       content.Category,
		Duration:       content.Duration,
		Resolution:     content.Resolution,
		FileSize:       content.FileSize,
		Format:         content.Format,
		Quality:        content.Quality,
		ApprovalStatus: content.ApprovalStatus,
	}
	db := c.data.db
	if err := db.Create(&detail).Error; err != nil {
		c.log.Errorf("content create error = %v", err)
		return err
	}
	return nil
}

func (c *contentRepo) Update(ctx context.Context, id int64, content *biz.Content) error {
	db := c.data.db
	detail := ContentDetail{
		ContentID:      content.ContentID,
		Title:          content.Title,
		Description:    content.Description,
		Author:         content.Author,
		VideoURL:       content.VideoURL,
		Thumbnail:      content.Thumbnail,
		Category:       content.Category,
		Duration:       content.Duration,
		Resolution:     content.Resolution,
		FileSize:       content.FileSize,
		Format:         content.Format,
		Quality:        content.Quality,
		ApprovalStatus: content.ApprovalStatus,
	}
	if err := db.Where("id = ?", id).
		Updates(&detail).Error; err != nil {
		c.log.WithContext(ctx).Errorf("content update error = %v", err)
		return err
	}
	return nil
}

func (c *contentRepo) IsExist(ctx context.Context, id int64) (bool, error) {
	db := c.data.db
	var detail ContentDetail
	err := db.Where("id = ?", id).First(&detail).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		c.log.WithContext(ctx).Errorf("ContentDao isExist = [%v]", err)
		return false, err
	}
	return true, nil
}

func (c *contentRepo) Delete(ctx context.Context, id int64) error {
	db := c.data.db
	// 删除索引信息
	err := db.Where("id = ?", id).
		Delete(&ContentDetail{}).Error
	if err != nil {
		c.log.WithContext(ctx).Errorf("content delete error = %v", err)
		return err
	}
	return nil
}

func (c *contentRepo) Find(ctx context.Context, params *biz.FindParams) ([]*biz.Content, int64, error) {
	// 构造查询条件
	query := c.data.db.Model(&ContentDetail{})
	if params.ID != 0 {
		query = query.Where("id = ?", params.ID)
	}
	if params.Author != "" {
		query = query.Where("author = ?", params.Author)
	}
	if params.Title != "" {
		query = query.Where("title = ?", params.Title)
	}
	// 总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	//设置默认页大小
	var page, pageSize = 1, 10
	if params.Page > 0 {
		page = int(params.Page)
	}
	if params.PageSize > 0 {
		pageSize = int(params.PageSize)
	}
	offset := (page - 1) * pageSize
	//进行数据库查找
	var results []*ContentDetail
	if err := query.Offset(offset).Limit(pageSize).
		Find(&results).Error; err != nil {
		c.log.WithContext(ctx).Errorf("content find error = %v", err)
		return nil, 0, err
	}
	var contents []*biz.Content
	//将数据库查找的结构，映射到biz.Content定义的结构
	for _, r := range results {
		contents = append(contents, &biz.Content{
			ID:             r.ID,
			Title:          r.Title,
			VideoURL:       r.VideoURL,
			Author:         r.Author,
			Description:    r.Description,
			Thumbnail:      r.Thumbnail,
			Category:       r.Category,
			Duration:       r.Duration,
			Resolution:     r.Resolution,
			FileSize:       r.FileSize,
			Format:         r.Format,
			Quality:        r.Quality,
			ApprovalStatus: r.ApprovalStatus,
			UpdatedAt:      r.UpdatedAt,
			CreatedAt:      r.CreatedAt,
		})
	}
	return contents, total, nil
}
