package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"singo/serializer"
	"singo/service"
)

// BusinessRegister 商家注册接口
func BusinessRegister(c *gin.Context) {
	var service service.BusinessRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// BusinessLogin 商家登录接口
func BusinessLogin(c *gin.Context) {
	var service service.BusinessLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// BusinessMe 商家详情
func BusinessMe(c *gin.Context) {
	business := CurrentBusiness(c)
	res := serializer.BuildBusinessResponse(*business)
	c.JSON(200, res)
}

// BusinessLogout 商家登出
func BusinessLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
