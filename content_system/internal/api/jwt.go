package api

import (
	"content_system/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//const (
//	SessionKey = "session_id" //鉴权的id，前端需要在Headers中加入这个参数session_id
//)

// jwt中间件,只验证jwt。(生成jwt放到login_handle的登录里面)
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader(SessionKey) //从请求头，拿到token
		if tokenHeader == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, "jwt 为空")
			return
		}
		//如果不为空，则进行验证
		key, err := jwt.CheckToken(tokenHeader)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, "jwt 鉴权失败")
			return
		}
		//信息上，鉴权成功
		//还需要判断token是否过期
		if time.Now().Unix() > key.ExpiresAt.Unix() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "jwt 鉴权过期")
			return
		}

		//c.Set("username", key.Username)
		c.Next()
	}
}
