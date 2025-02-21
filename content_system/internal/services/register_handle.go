package services

import (
	"content_system/internal/dao"
	"content_system/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// 模型的绑定与验证：绑定：【处理我们自定义的数据结构体】；字段验证：【不能为空，手机号规则，邮箱规则,长度】
// 前端请求的结构体数据集【注册账号，前端要传给我：用户名、密码、名称这3个信息】
type RegisterReq struct {
	UserID   string `json:"user_id" binding:"required"` //定义前端传过来，必须要包含该字段(required)
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
}

// 响应结构体
type RegisterRes struct {
	Message string `json:"message" binding:"required"` //定义返回内容，返回一个json消息
}

func (cms *CmsAPP) Register(c *gin.Context) {
	//初始化前端的返回结构实例
	var req RegisterReq

	//要求前端按照HelloReq结构体中定义的那样，如果不按照，则会报错
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//如果前端正确发送请求，则执行下面的程序，返回响应数据
	fmt.Printf("register info = %+v \n", req) //%v是只打印【值】，%+v是打印【字段+值】

	//密码要实现加密
	hashedPassword, err := encryptPassword(req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	fmt.Printf("hashed password = [%s]\n", hashedPassword)

	//初始化dao层的实例，用dao层的方法，实现功能逻辑
	accountDao := dao.NewAccountDao(cms.db)

	//账号校验（数据库中存在该账号的话，要提示注册错误）
	isExist, err := accountDao.IsExist(req.UserID)
	if err != nil { //注册发生错误
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if isExist { //如果账号存在的话
		c.JSON(http.StatusBadRequest, gin.H{"error": "账号已存在"})
		return
	} //如果不存在，则执行下面的内容

	//账号信息持久化(写入数据库)
	if err := accountDao.Create(model.Account{ //填写Account所有字段
		UserID:   req.UserID,
		Password: hashedPassword,
		Nickname: req.Nickname,
	}); err != nil { //如果发生错误
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//回包
	c.JSON(http.StatusOK, gin.H{ //注册成功
		"msg": "ok",
		"data": &RegisterRes{
			Message: "注册成功",
		},
	})
}

// 密码加密函数
func encryptPassword(password string) (string, error) {
	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if error != nil {
		fmt.Println("Error hashing password:", error)
		return "", error
	}
	return string(hashedPassword), nil
}
