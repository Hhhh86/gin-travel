package service

import (
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
	"time"
)

// ProductInfo 景点信息
type ProductInfoService struct {
	SID         int    `json:"s_id"`         //景点id
	BID         int    `json:"b_id"`         //公司id
	Number      int    `json:"number"`       //人数
	Price       int    `json:"price"`        //套餐价格
	CompanyName string `json:"company_name"` //公司名称
	Name        string `json:"name"`         //套餐名称
	Describe    string `json:"describe"`     //描述
	StartTime   string `json:"start_time"`   //出发时间
	EndTime     string `json:"end_time"`     //结束时间
}

//CreateProduct 新建景点旅游产品
func (service *ProductInfoService) CreateProduct(c *gin.Context) serializer.Response {
	//查询景点是否存在
	var spot model.SpotInfo
	err := model.DB.First(&spot, service.SID).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "景点不存在",
			Error: err.Error(),
		}
	}
	//根据上下文查询公司名称
	if business, _ := c.Get("business"); business != nil {
		if b, ok := business.(*model.Business); ok {
			service.BID = int(b.ID)
			service.CompanyName = b.CompanyName
		}

	}
	var timeLayoutStr = "2006-01-02 15:04:05"                  //go中的时间格式化必须是这个时间
	start, _ := time.Parse(timeLayoutStr, service.StartTime) //string转time
	end, _ := time.Parse(timeLayoutStr, service.EndTime)

	info := model.ProductInfo{
		SID:         service.SID,
		BID:         service.BID,
		Number:      service.Number,
		Price:       service.Price,
		CompanyName: service.CompanyName,
		Name:        service.Name,
		Describe:    service.Describe,
		StartTime:   start,
		EndTime:     end,
	}
	if err := model.DB.Create(&info).Error; err != nil {
		return serializer.ParamErr("新建产品信息失败", err)
	}

	return serializer.SuccessResponse()

}
