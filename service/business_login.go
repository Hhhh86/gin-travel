package service

import (
	"singo/model"
	"singo/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// BusinessLoginService 管理商家登录的服务
type BusinessLoginService struct {
	CompanyName string `form:"company_name" json:"company_name" binding:"required,min=0,max=30"`
	Password    string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// setSession 设置session
func (service *BusinessLoginService) setSession(c *gin.Context, business model.Business) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("business_id", business.ID)
	s.Save()
}

// Login 商家登录函数
func (service *BusinessLoginService) Login(c *gin.Context) serializer.Response {
	var business model.Business

	if err := model.DB.Where("company_name = ?", service.CompanyName).First(&business).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	if business.CheckPassword(service.Password) == false {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	// 设置session
	service.setSession(c, business)

	return serializer.BuildBusinessResponse(business)
}
