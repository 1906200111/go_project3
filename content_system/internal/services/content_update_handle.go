package services

import (
	"github.com/gin-gonic/gin"
	"go_project2/internal/dao"
	"go_project2/internal/model"
	"net/http"
	"time"
)

// 前端更新，需要传的内容结构。id是必须传的，需要知道对哪条内容进行更新操作
type ContentUpdateReq struct {
	ID             int           `json:"id" binding:"required"` // 内容标题
	Title          string        `json:"title"`                 // 内容标题
	VideoURL       string        `json:"video_url"`             // 视频播放URL
	Author         string        `json:"author"`                // 作者
	Description    string        `json:"description"`           // 内容描述
	Thumbnail      string        `json:"thumbnail"`             // 封面图URL
	Category       string        `json:"category"`              // 内容分类
	Duration       time.Duration `json:"duration"`              // 内容时长
	Resolution     string        `json:"resolution"`            // 分辨率 如720p、1080p
	FileSize       int64         `json:"fileSize"`              // 文件大小
	Format         string        `json:"format"`                // 文件格式 如MP4、AVI
	Quality        int           `json:"quality"`               // 视频质量 1-高清 2-标清
	ApprovalStatus int           `json:"approval_status"`       // 审核状态 1-审核中 2-审核通过 3-审核不通过
}

// 后端返回的响应结构
type ContentUpdateRsp struct {
	Message string `json:"message"`
}

func (c *CmsAPP) ContentUpdate(ctx *gin.Context) {
	var req ContentUpdateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//实例化dao层的实例
	contentDao := dao.NewContentDao(c.db)
	//查看该content是否存在，存在的话，才能进行更新
	ok, err := contentDao.IsExist(req.ID)
	if err != nil { //出现错误
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok { //该内容不存在，更新操作就无法实现
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "内容不存在，无法更新"})
		return
	}
	//如果存在这个内容，则进行更新操作
	if err := contentDao.Update(req.ID, model.ContentDetail{
		Title:          req.Title,
		Description:    req.Description,
		Author:         req.Author,
		VideoURL:       req.VideoURL,
		Thumbnail:      req.Thumbnail,
		Category:       req.Category,
		Duration:       req.Duration,
		Resolution:     req.Resolution,
		FileSize:       req.FileSize,
		Format:         req.Format,
		Quality:        req.Quality,
		ApprovalStatus: req.ApprovalStatus,
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &ContentUpdateRsp{
			Message: "content更新ok",
		},
	})
}
