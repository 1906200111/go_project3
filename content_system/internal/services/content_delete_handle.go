package services

import (
	"github.com/gin-gonic/gin"
	"go_project2/internal/dao"
	"net/http"
)

// 前端需要传的参数：id是必须的，因为需要知道删除的是哪个数据
type ContentDeleteReq struct {
	ID int `json:"id" binding:"required"` // 内容ID
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
	//实例化dao层的实例
	contentDao := dao.NewContentDao(c.db)
	//先判断这个数据是否存在
	ok, err := contentDao.IsExist(req.ID)
	//出错
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"IsExist error": err.Error()})
		return
	} //内容不存在
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "内容本身不存在，无法删除"})
		return
	} //内容存在的话，则进行删除
	if err := contentDao.Delete(req.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"delete error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &ContentDeleteRsp{
			Message: "content删除ok",
		},
	})
}
