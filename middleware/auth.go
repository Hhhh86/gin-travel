package middleware

import (
	"singo/model"
	"singo/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			user, err := model.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired(open bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !open{
			c.Next()
			return
		}
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}
		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}

// CurrentBusiness 获取登录商家
func CurrentBusiness() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		bid := session.Get("business_id")
		if bid != nil {
			business, err := model.GetBusiness(bid)
			if err == nil {
				c.Set("business", &business)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthBusinessRequired(open bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !open{
			c.Next()
			return
		}
		if business, _ := c.Get("business"); business != nil {
			if _, ok := business.(*model.Business); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}
