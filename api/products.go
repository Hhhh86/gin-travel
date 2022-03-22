package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// CreateProduct 新建产品
func CreateProduct(c *gin.Context) {
	var service service.BusinessLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}