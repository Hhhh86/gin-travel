package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// CreateProduct 新建产品
func CreateProduct(c *gin.Context) {
	var service service.CreateProductService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateProduct(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteProduct 删除产品
func DeleteProduct(c *gin.Context) {
	var service service.DeleteProductService
	res := service.DeleteProduct(c, c.Param("id"))
	c.JSON(200, res)
}

// UpdateSpot 更新景点信息
func UpdateProduct(c *gin.Context) {
	var service service.UpdateProductService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateProduct(c, c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AllProduct 获取全部产品信息
func AllProduct(c *gin.Context) {
	var service service.ListProductService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListProduct(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AllProduct 获取全部产品信息
func GetProductByID(c *gin.Context) {
	var service service.GetProductService
	res := service.GetProductByID(c, c.Param("id"))
	c.JSON(200, res)
}


// CollectProduct 收藏产品
func CollectProduct(c *gin.Context) {
	var service service.CollectProductService
	res := service.CollectProduct(c, c.Param("pid"))
	c.JSON(200, res)
}

// LikeProduct 点赞产品
func LikeProduct(c *gin.Context) {
	var service service.LikeProductService
	res := service.LikeProduct(c, c.Param("pid"))
	c.JSON(200, res)
}

// LikeProduct 点赞产品
func DisLikeProduct(c *gin.Context) {
	var service service.LikeProductService
	res := service.DisLikeProduct(c, c.Param("pid"))
	c.JSON(200, res)
}
