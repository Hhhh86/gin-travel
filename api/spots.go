package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// CreateSpot 新建景点信息
func CreateSpot(c *gin.Context) {
	var service service.CreateSpotService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateSpot(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteSpot 删除景点信息
func DeleteSpot(c *gin.Context) {
	service := service.DeleteSpotService{}
	res := service.DeleteSpot(c, c.Param("id"))
	c.JSON(200, res)
}

// UpdateSpot 更新景点信息
func UpdateSpot(c *gin.Context) {
	var service service.UpdateSpotService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateSpot(c, c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AllSpot 获取全部景点信息
func AllSpot(c *gin.Context) {
	var service service.ListSpotService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListSpot(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// CollectSpot 收藏景点
func CollectSpot(c *gin.Context) {
	var service service.CollectSpotService
	res := service.CollectSpot(c, c.Param("sid"))
	c.JSON(200, res)
}
