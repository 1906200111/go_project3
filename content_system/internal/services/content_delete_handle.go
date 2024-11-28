package services

import (
	"content_system/internal/api/operate"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 前端需要传的参数：id是必须的，因为需要知道删除的是哪个数据
type ContentDeleteReq struct {
	ID int64 `json:"id" binding:"required"` // 内容ID
}

// 后端返回的结构
type ContentDeleteRsp struct {
	Message string `json:"message"`
}

func (c *CmsAPP) ContentDelete(ctx *gin.Context) {
	var req ContentDeleteReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//下面不走，直接db的方法，走的是grpc的方法。【内容网关功能很干净了，不走db的操作，转发给grpc去执行操作】
	rsp, err := c.operateAppClient.DeleteContent(ctx, &operate.DeleteContentReq{Id: req.ID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"IsExist error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": rsp,
	})
}
