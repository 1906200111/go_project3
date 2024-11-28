package api

import (
	"content_system/internal/utils"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
)

const (
	SessionKey = "session_id" //鉴权的id，前端需要在Headers中加入这个参数session_id
)

// session鉴权
type SessionAuth struct {
	rdb *redis.Client
}

func NewSessionAuth() *SessionAuth {
	s := &SessionAuth{} //创建中间件实例
	connRdb(s)          //给实例加上redis连接的方法
	return s
}

func (s *SessionAuth) Auth(c *gin.Context) {
	sessionID := c.GetHeader(SessionKey)

	if sessionID == "" { //前端返回session为空，拦截，并且返回一段话
		c.AbortWithStatusJSON(http.StatusForbidden, "session 为空")
	}
	authKey := utils.GetAuthKey(sessionID)
	loginTime, err := s.rdb.Get(c, authKey).Result()
	if err != nil && err != redis.Nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "session auth error")
	}
	if loginTime == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "session 鉴权失败")
	}
	c.Next()
}

func connRdb(s *SessionAuth) {
	//redis-cli
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	s.rdb = rdb
}
