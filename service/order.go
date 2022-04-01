package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"singo/model"
	"singo/serializer"
)

const (
	UNDONE = 1
	DONE   = 2
	CANCEL = 3
)

// CreateOrderService 景点信息
type CreateOrderService struct {
	PID    int `json:"p_id"`   //产品id
	Number int `json:"number"` //人数
}

//CreateOrder 新建旅游产品订单
func (service *CreateOrderService) CreateOrder(c *gin.Context) serializer.Response {
	//查询产品是否存在
	var product model.ProductInfo
	err := model.DB.First(&product, service.PID).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "产品不存在",
			Error: err.Error(),
		}
	}
	var user model.User
	//根据上下文查询用户身份
	if c_user, _ := c.Get("user"); c_user != nil {
		if u, ok := c_user.(*model.User); ok {
			err := model.DB.Where("user_name=?", u.UserName).First(&user).Error
			if err != nil {
				return serializer.Response{
					Code:  404,
					Msg:   "用户不存在",
					Error: err.Error(),
				}
			}
		}
	}

	info := model.OrderInfo{
		PID:      service.PID,
		Number:   service.Number,
		Status:   UNDONE, //未完成
		UserName: user.UserName,
	}
	if err := model.DB.Create(&info).Error; err != nil {
		return serializer.ParamErr("新建订单失败", err)
	}

	return serializer.SuccessResponse()

}

// OrderAllInfo 查看景点信息
type RetrieveOrderService struct {
}

func (service *RetrieveOrderService) RetrieveAllOrder(c *gin.Context) serializer.Response {
	//根据上下文查询用户身份
	var order []model.OrderInfo
	if c_user, _ := c.Get("user"); c_user != nil {
		if u, ok := c_user.(*model.User); ok {
			err := model.DB.Find(&order).Where("user_name=?", u.UserName).Error
			if err != nil && err == gorm.ErrRecordNotFound {
				return serializer.Response{
					Code:  404,
					Msg:   "订单不存在",
					Error: err.Error(),
				}
				return serializer.Response{
					Code:  500,
					Msg:   "内部出现问题，请稍后尝试",
					Error: err.Error(),
				}
			}
		}
	}
	return serializer.BuildAllOrderResponse(order)
}

func (service *RetrieveOrderService) RetrieveOneOrder(c *gin.Context, id string) serializer.Response {
	//根据上下文查询用户身份
	var order model.OrderInfo
	var product model.ProductInfo
	if c_user, _ := c.Get("user"); c_user != nil {
		if u, ok := c_user.(*model.User); ok {
			err := model.DB.Where("id = ? and user_name= ? ", id, u.UserName).First(&order).Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					return serializer.Response{
						Code:  404,
						Msg:   "订单不存在",
						Error: err.Error(),
					}
				}
				return serializer.Response{
					Code:  500,
					Msg:   "内部出现问题",
					Error: err.Error(),
				}
			}

			err = model.DB.First(&product, order.PID).Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					return serializer.Response{
						Code:  404,
						Msg:   "产品不存在",
						Error: err.Error(),
					}
				}
				return serializer.Response{
					Code:  500,
					Msg:   "内部出现问题",
					Error: err.Error(),
				}
			}

		}
	}
	return serializer.BuildOneOrderResponse(order, product)
}

// OrderAllInfo 查看景点信息
type CancelOrderService struct {
}

func (service *CancelOrderService) CancelOrder(c *gin.Context, id string) serializer.Response {
	//根据上下文查询用户身份
	if c_user, _ := c.Get("user"); c_user != nil {
		if u, ok := c_user.(*model.User); ok {
			err := model.DB.Model(&model.OrderInfo{}).Where("id = ? and user_name = ?", id, u.UserName).Update("status", CANCEL).Error
			if err != nil {
				return serializer.Response{
					Code:  500,
					Msg:   "内部出现问题，请稍后尝试",
					Error: err.Error(),
				}
			}
		}
	}
	return serializer.SuccessResponse()
}

// OrderAllInfo 查看景点信息
type DeleteOrderService struct {
}

func (service *DeleteOrderService) DeleteOrder(c *gin.Context, id string) serializer.Response {
	var order model.OrderInfo
	//根据上下文查询用户身份
	if c_user, _ := c.Get("user"); c_user != nil {
		if u, ok := c_user.(*model.User); ok {
			//查询是否有该订单
			err := model.DB.First(&order).Where("id = ? and user_name = ?", id, u.UserName).Error
			if err != nil {
				return serializer.Response{
					Code:  404,
					Msg:   "订单不存在",
					Error: err.Error(),
				}
			}

			err = model.DB.Delete(&order).Error
			if err != nil {
				return serializer.Response{
					Code:  500,
					Msg:   "删除订单错误",
					Error: err.Error(),
				}
			}

		}
	}
	return serializer.SuccessResponse()
}

// CreateOrderComment 新建订单评论
type CreateOrderCommentService struct {
	OID     int    `json:"o_id"`    //订单id
	Content string `json:"content"` //评论内容
}

//CreateOrder 新建旅游产品订单
func (service *CreateOrderCommentService) CreateOrderComment(c *gin.Context) serializer.Response {
	//查询订单是否存在且为已完成
	var order model.OrderInfo
	err := model.DB.Where("id = ? and status = ?", service.OID, DONE).First(&order).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "订单不存在或尚未完成",
			Error: err.Error(),
		}
	}
	var username string
	//根据上下文查询用户身份
	if c_user, _ := c.Get("user"); c_user != nil {
		if u, ok := c_user.(*model.User); ok {
			if u.UserName != order.UserName {
				return serializer.Response{
					Code:  400,
					Msg:   "身份不一致",
					Error: err.Error(),
				}
			}
			//查询该订单是否已经有该评价
			count := int64(0)
			model.DB.Model(&model.OrderComment{}).Where("o_id = ? and user_name = ?", service.OID, u.UserName).Count(&count)
			username = u.UserName
			if count > 0 {
				return serializer.Response{
					Code:  400,
					Msg:   "该订单已评价",
				}
			}
		}
	}

	comment := model.OrderComment{
		OID:      service.OID,
		Content:  service.Content,
		UserName: username,
	}
	if err := model.DB.Create(&comment).Error; err != nil {
		return serializer.ParamErr("新建评价失败", err)
	}

	return serializer.SuccessResponse()

}
