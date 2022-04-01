package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// CreateOrder 新建订单
func CreateOrder(c *gin.Context) {
	var service service.CreateOrderService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateOrder(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// RetrieveAllOrder 查看全部订单
func RetrieveAllOrder(c *gin.Context) {
	var service service.RetrieveOrderService
	res := service.RetrieveAllOrder(c)
	c.JSON(200, res)
}

// RetrieveOneOrder 查看单一订单
func RetrieveOneOrder(c *gin.Context) {
	var service service.RetrieveOrderService
	res := service.RetrieveOneOrder(c, c.Param("id"))
	c.JSON(200, res)
}

// CancelOrder 取消订单
func CancelOrder(c *gin.Context) {
	service := service.CancelOrderService{}
	res := service.CancelOrder(c, c.Param("id"))
	c.JSON(200, res)
}

// DeleteOrder 删除订单
func DeleteOrder(c *gin.Context) {
	service := service.DeleteOrderService{}
	res := service.DeleteOrder(c, c.Param("id"))
	c.JSON(200, res)
}


// EvaluateOrder 评价订单
func EvaluateOrder(c *gin.Context) {
	var service service.CreateOrderCommentService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateOrderComment(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
