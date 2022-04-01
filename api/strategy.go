package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// CreateStrategy 新建攻略
func CreateStrategy(c *gin.Context) {
	var service service.CreateStrategyService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateStrategy(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteStrategy 删除攻略
func DeleteStrategy(c *gin.Context) {
	var service service.DeleteStrategyService
		res := service.DeleteStrategy(c,c.Param("sid"))
		c.JSON(200, res)
}

// CreateStrategyComment 评论攻略
func CreateStrategyComment(c *gin.Context) {
	var service service.CreateStrategyCommentService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateComment(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteStrategyComment 删除攻略评论
func DeleteStrategyComment(c *gin.Context) {
	var service service.DeleteStrategyCommentService
	res := service.DeleteComment(c,c.Param("stid"))
	c.JSON(200, res)
}


// LikeStrategy 点赞攻略
func LikeStrategy(c *gin.Context) {
	var service service.LikeStrategyService
	res := service.LikeStrategy(c, c.Param("stid"))
	c.JSON(200, res)
}

// DisLikeStrategy 取消点赞攻略
func DisLikeStrategy(c *gin.Context) {
	var service service.LikeStrategyService
	res := service.DisLikeStrategy(c, c.Param("stid"))
	c.JSON(200, res)
}
