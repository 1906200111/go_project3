package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 模型的绑定与验证：绑定：【处理我们自定义的数据结构体】；字段验证：【不能为空，手机号规则，邮箱规则,长度】
// 请求结构体
type HelloReq struct {
	Name string `json:"name" binding:"required"` //定义前端传过来，必须要包含name字段(required)
}

// 响应结构体
type HelloRes struct {
	Message string `json:"message" binding:"required"` //定义前端传过来，必须要包含name字段(required)
}

func (cms *CmsAPP) Hello(c *gin.Context) {
	var req HelloReq
	//要求前端按照HelloReq结构体中定义的那样，如果不按照，则会报错
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//如果前端正确发送请求，则执行下面的程序，返回响应数据
	c.JSON(http.StatusOK, gin.H{
		"name": req.Name,
		"data": &HelloRes{
			Message: "hello" + req.Name,
		},
	})
}
