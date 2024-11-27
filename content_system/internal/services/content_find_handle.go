package services

import (
	"github.com/gin-gonic/gin"
	"go_project2/internal/dao"
	"net/http"
	"time"
)

type Content struct {
	ID             int           `json:"id"`                        // 内容ID
	Title          string        `json:"title"`                     // 内容标题
	VideoURL       string        `json:"video_url" `                // 视频播放URL
	Author         string        `json:"author" binding:"required"` // 作者
	Description    string        `json:"description"`               // 内容描述
	Thumbnail      string        `json:"thumbnail"`                 // 封面图URL
	Category       string        `json:"category"`                  // 内容分类
	Duration       time.Duration `json:"duration"`                  // 内容时长
	Resolution     string        `json:"resolution"`                // 分辨率 如720p、1080p
	FileSize       int64         `json:"fileSize"`                  // 文件大小
	Format         string        `json:"format"`                    // 文件格式 如MP4、AVI
	Quality        int           `json:"quality"`                   // 视频质量 1-高清 2-标清
	ApprovalStatus int           `json:"approval_status"`
}

// 前端的请求数据结构
type ContentFindReq struct {
	ID       int    `json:"id"`        // 内容ID
	Author   string `json:"author"`    // 作者
	Title    string `json:"title"`     // 标题
	Page     int    `json:"page"`      // 页
	PageSize int    `json:"page_size"` // 页大小
}

// 后端响应的数据结构
type ContentFindRsp struct {
	Message  string    `json:"message"`
	Contents []Content `json:"contents"` //包含多个content的实例，用切片返回
	Total    int64     `json:"total"`
}

func (c *CmsAPP) ContentFind(ctx *gin.Context) {
	var req ContentFindReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//定义dao层的实例
	contentDao := dao.NewContentDao(c.db)
	//调用查询Find方法
	contentList, total, err := contentDao.Find(&dao.FindParams{ //把前端的查询组合参数，传进去
		ID:       req.ID,
		Author:   req.Author,
		Title:    req.Title,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil { //如果出错
		ctx.JSON(http.StatusInternalServerError, gin.H{"Find content error": err.Error()})
		return
	}
	//成功查询到数据
	contents := make([]Content, 0, len(contentList))
	for _, content := range contentList {
		contents = append(contents, Content{
			ID:             content.ID,
			Title:          content.Title,
			VideoURL:       content.VideoURL,
			Author:         content.Author,
			Description:    content.Description,
			Thumbnail:      content.Thumbnail,
			Category:       content.Category,
			Duration:       content.Duration,
			Resolution:     content.Resolution,
			FileSize:       content.FileSize,
			Format:         content.Format,
			Quality:        content.Quality,
			ApprovalStatus: content.ApprovalStatus,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &ContentFindRsp{
			Message:  "content查询ok",
			Contents: contents,
			Total:    total,
		},
	})
}
