package service

import (
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
)

// CreateSpotService 新建景点信息结构体
type CreateSpotService struct {
	AgeScope     int    `json:"age_scope"`  //1Children、2Youth、3Midlife、4Aged、5all
	PlaceType    int    `json:"place_type"` //1自然生态类、2历史文化类、3现代游乐类、4产业融合类、5其他类。
	Price        int    `json:"price"`      //门票价格
	Name         string `json:"name" binding:"required,min=1,max=30"`
	Address      string `json:"address" binding:"required,min=3,max=99"`
	Introduction string `json:"introduction" binding:"required,min=0,max=99"`
	Tips         string `json:"tips" binding:"required,min=0,max=99"`
	OpenTime     string `json:"open_time" binding:"required,min=0,max=20"`
}

// CreateSpot 新建景点信息
func (service *CreateSpotService) CreateSpot(c *gin.Context) serializer.Response {
	// 新建景点
	info := model.SpotInfo{
		Name:         service.Name,
		Address:      service.Address,
		AgeScope:     service.AgeScope,
		Introduction: service.Introduction,
		Tips:         service.Tips,
		OpenTime:     service.OpenTime,
		PlaceType:    service.PlaceType,
		Price:        service.Price,
	}
	if err := model.DB.Create(&info).Error; err != nil {
		return serializer.ParamErr("新建景点信息失败", err)
	}

	return serializer.SuccessResponse()
}

// UpdateSpotService 删除景点的服务
type DeleteSpotService struct {
}

// DeleteSpot 删除景点信息
func (service *DeleteSpotService) DeleteSpot(c *gin.Context,id string) serializer.Response {
	var info model.SpotInfo
	err := model.DB.First(&info, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "景点不存在",
			Error: err.Error(),
		}
	}
	err = model.DB.Delete(&info).Error
	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "景点删除失败",
			Error: err.Error(),
		}
	}

	return serializer.SuccessResponse()
}

// CreateSpotService 新建景点信息结构体
type UpdateSpotService struct {
	AgeScope     int    `json:"age_scope"`  //1Children、2Youth、3Midlife、4Aged、5all
	PlaceType    int    `json:"place_type"` //1自然生态类、2历史文化类、3现代游乐类、4产业融合类、5其他类。
	Price        int    `json:"price"`      //门票价格
	Address      string `json:"address" binding:"required,min=3,max=99"`
	Introduction string `json:"introduction" binding:"required,min=0,max=99"`
	Tips         string `json:"tips" binding:"required,min=0,max=99"`
	OpenTime     string `json:"open_time" binding:"required,min=0,max=20"`
}

// CreateSpot 新建景点信息
func (service *UpdateSpotService) UpdateSpot(c *gin.Context,id string) serializer.Response {
	// 查找景点
	var info model.SpotInfo
	err := model.DB.First(&info, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "景点不存在",
			Error: err.Error(),
		}
	}
	info.Address = service.Address
	info.AgeScope = service.AgeScope
	info.Introduction = service.Introduction
	info.Tips = service.Tips
	info.OpenTime = service.OpenTime
	info.PlaceType = service.PlaceType
	info.Price = service.Price

	err = model.DB.Save(&info).Error
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "景点保存失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.SuccessResponse(),
	}
}

// ListSpotService 全部景点信息结构体
type ListSpotService struct {
	Start int `json:"start"`
	Limit int `json:"limit"`
}

// List 景点列表
func (service *ListSpotService) ListSpot(c *gin.Context) serializer.Response {
	infos := []model.SpotInfo{}
	//var total int64

	if service.Limit == 0 {
		service.Limit = 6
	}

	//if err := model.DB.Model(model.SpotInfo{}).Count(&total).Error; err != nil {
	//	return serializer.Response{
	//		Code: 50000,
	//		Msg:    "数据库连接错误",
	//		Error:  err.Error(),
	//	}
	//}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&infos).Error; err != nil {
		return serializer.Response{
			Code: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildSpotResponse(infos)
}
