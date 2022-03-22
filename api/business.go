package api

import (
	"github.com/gin-gonic/gin"
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

