package service

import (
	"singo/model"
	"singo/serializer"
)

// BusinessRegisterService 管理用户注册服务
type BusinessRegisterService struct {
	CompanyName     string `form:"company_name" json:"company_name" binding:"required,min=2,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// valid 验证表单
func (service *BusinessRegisterService) valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}

	count := int64(0)
	model.DB.Model(&model.Business{}).Where("company_name = ?", service.CompanyName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "公司名称已被占用",
		}
	}

	return nil
}

// Register 用户注册
func (service *BusinessRegisterService) Register() serializer.Response {
	business := model.Business{
		CompanyName: service.CompanyName,
		Status:      model.Active,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := business.SetPassword(service.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	if err := model.DB.Create(&business).Error; err != nil {
		return serializer.ParamErr("注册失败", err)
	}

	return serializer.BuildBusinessResponse(business)
}
