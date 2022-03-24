package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// Session 初始化session
func Session(secret string) gin.HandlerFunc {
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "123456", []byte(secret))
	return sessions.Sessions("gin-session", store)
}
