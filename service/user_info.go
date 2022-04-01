package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"singo/model"
	"singo/serializer"
)

// CreateInfoService 管理用户登录的服务
type CreateInfoService struct {
	Sex            int    `json:"sex"`              //1女 2男
	PlaceType      int    `json:"place_type"`       //1自然生态类、2历史文化类、3现代游乐类、4产业融合类、5其他类。
	AgeScope       int    `json:"age_scope"`        //0Children、1Youth、2Midlife，3Aged
	ExpectMinPrice int    `json:"expect_min_price"` //期望最小价格
	ExpectMaxPrice int    `json:"expect_max_price"` //期望最大价格
	UserName       string `json:"user_name"`
}

// CreateSpot 新建用户信息
func (service *CreateInfoService) CreateInfo(c *gin.Context) serializer.Response {
	var user model.User
	//根据上下文查询用户身份
	if c_user, _ := c.Get("user"); c_user != nil {
		if u, ok := c_user.(*model.User); ok {
			if u.UserName != service.UserName {
				return serializer.Response{
					Code: 405,
					Msg:  "用户身份不一致",
				}
			}
			err := model.DB.First(&user).Where("user_name=?", service.UserName).Error
			if err != nil {
				return serializer.Response{
					Code:  404,
					Msg:   "用户不存在",
					Error: err.Error(),
				}
			}
		}
	}
	var count int64
	model.DB.Model(&model.UserInfo{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return serializer.Response{
			Code: 40001,
			Msg:  "用户名信息已经存在",
		}
	}

	// 新建用户信息
	info := model.UserInfo{
		Sex:            service.Sex,
		PlaceType:      service.PlaceType,
		AgeScope:       service.AgeScope,
		ExpectMinPrice: service.ExpectMinPrice,
		ExpectMaxPrice: service.ExpectMaxPrice,
		UserName:       service.UserName,
		NickName:       user.NickName,
		Avatar:         user.Avatar,
	}
	if err := model.DB.Create(&info).Error; err != nil {
		return serializer.ParamErr("新建用户信息失败", err)
	}

	return serializer.SuccessResponse()
}

// CreateInfoService 管理用户登录的服务
type DeleteInfoService struct{}

func (service *DeleteInfoService) DeleteInfo(c *gin.Context) serializer.Response {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			var info model.UserInfo
			err := model.DB.Where("user_name = ?", u.UserName).Delete(&info).Error
			if err != nil {
				return serializer.Response{
					Code:  50000,
					Msg:   "用户信息删除失败",
					Error: err.Error(),
				}
			}
		}
	}
	return serializer.SuccessResponse()
}

// RetrieveInfoService 用户查看信息
type RetrieveInfoService struct{}

func (service *RetrieveInfoService) RetrieveInfo(c *gin.Context) serializer.Response {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			var info model.UserInfo
			err := model.DB.Where("user_name = ?", u.UserName).First(&info).Error
			if err != nil {
				return serializer.Response{
					Code:  40001,
					Msg:   "用户信息为空",
					Error: err.Error(),
				}
			}
			return serializer.BuildUserInfoResponse(info)
		}
		return serializer.Response{
			Code: 40001,
			Msg:  "用户信息为空",
		}
	}
	return serializer.SuccessResponse()
}

// UpdateInfoService 管理用户登录的服务
type UpdateInfoService struct {
	Sex            int    `json:"sex"`              //1女 2男
	PlaceType      int    `json:"place_type"`       //1自然生态类、2历史文化类、3现代游乐类、4产业融合类、5其他类。
	AgeScope       int    `json:"age_scope"`        //0Children、1Youth、2Midlife，3Aged
	ExpectMinPrice int    `json:"expect_min_price"` //期望最小价格
	ExpectMaxPrice int    `json:"expect_max_price"` //期望最大价格
	UserName       string `json:"user_name"`
}

// CreateSpot 新建用户信息
func (service *UpdateInfoService) UpdateInfo(c *gin.Context) serializer.Response {
	var user model.User
	//根据上下文查询用户身份
	if c_user, _ := c.Get("user"); c_user != nil {
		if u, ok := c_user.(*model.User); ok {
			if u.UserName != service.UserName {
				return serializer.Response{
					Code: 405,
					Msg:  "用户身份不一致",
				}
			}
			err := model.DB.First(&user).Where("user_name=?", service.UserName).Error
			if err != nil {
				return serializer.Response{
					Code:  404,
					Msg:   "用户不存在",
					Error: err.Error(),
				}
			}
		}
	}
	var info model.UserInfo
	err := model.DB.Where("user_name = ?", service.UserName).First(&info).Error
	if err == gorm.ErrRecordNotFound {
		return serializer.Response{
			Code: 40001,
			Msg:  "该用户信息不存在，请先新增信息",
		}
	}
	// 更改用户信息
	info.Sex = service.Sex
	info.PlaceType = service.PlaceType
	info.AgeScope = service.AgeScope
	info.ExpectMinPrice = service.ExpectMinPrice
	info.ExpectMaxPrice = service.ExpectMaxPrice
	info.UserName = service.UserName

	if err := model.DB.Save(&info).Error; err != nil {
		return serializer.ParamErr("修改用户信息失败", err)
	}

	return serializer.SuccessResponse()
}
