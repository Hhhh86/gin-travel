package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"singo/cache"
	"singo/model"
	"singo/serializer"
	"time"
)

// ProductInfo 景点信息
type CreateProductService struct {
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
func (service *CreateProductService) CreateProduct(c *gin.Context) serializer.Response {
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

			var timeLayoutStr = "2006-01-02 15:04:05"                //go中的时间格式化必须是这个时间
			start, _ := time.Parse(timeLayoutStr, service.StartTime) //string转time
			end, _ := time.Parse(timeLayoutStr, service.EndTime)

			info := model.ProductInfo{
				SID:         service.SID,
				BID:         int(b.ID),
				CompanyName: b.CompanyName,
				Number:      service.Number,
				Price:       service.Price,
				Name:        service.Name,
				Describe:    service.Describe,
				StartTime:   start,
				EndTime:     end,
			}
			if err := model.DB.Create(&info).Error; err != nil {
				return serializer.ParamErr("新建产品信息失败", err)
			}
		}
	}

	return serializer.SuccessResponse()

}

// UpdateProductService 景点信息
type UpdateProductService struct {
	Number    int    `json:"number"`     //人数
	Price     int    `json:"price"`      //套餐价格
	Name      string `json:"name"`       //套餐名称
	Describe  string `json:"describe"`   //描述
	StartTime string `json:"start_time"` //出发时间
	EndTime   string `json:"end_time"`   //结束时间
}

//UpdateProduct 更新景点旅游产品
func (service *UpdateProductService) UpdateProduct(c *gin.Context, id string) serializer.Response {
	var info model.ProductInfo

	//根据上下文查询公司名称是否为该产品公司
	if business, _ := c.Get("business"); business != nil {
		if b, ok := business.(*model.Business); ok {
			//查询产品是否存在
			err := model.DB.First(&info).Where("id = ? and company_name = ?", id, b.CompanyName).Error
			if err != nil {
				return serializer.Response{
					Code:  404,
					Msg:   "产品不存在",
					Error: err.Error(),
				}
			}
		}

	}

	var timeLayoutStr = "2006-01-02 15:04:05"                //go中的时间格式化必须是这个时间
	start, _ := time.Parse(timeLayoutStr, service.StartTime) //string转time
	end, _ := time.Parse(timeLayoutStr, service.EndTime)

	info.Number = service.Number
	info.Price = service.Price
	info.Name = service.Name
	info.Describe = service.Describe
	info.StartTime = start
	info.EndTime = end

	err := model.DB.Save(&info).Error
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "产品保存失败",
			Error: err.Error(),
		}
	}

	return serializer.SuccessResponse()

}

// UpdateSpotService 删除景点的服务
type DeleteProductService struct {
}

//DeleteProduct 删除景点旅游产品
func (service *DeleteProductService) DeleteProduct(c *gin.Context, id string) serializer.Response {
	var info model.ProductInfo

	//根据上下文查询公司名称是否为该产品公司
	if business, _ := c.Get("business"); business != nil {
		if b, ok := business.(*model.Business); ok {
			//查询产品是否存在
			err := model.DB.First(&info).Where("id = ? and company_name = ?", id, b.CompanyName).Error
			if err != nil {
				return serializer.Response{
					Code:  404,
					Msg:   "产品不存在",
					Error: err.Error(),
				}
			}
		}

	}

	err := model.DB.Delete(&info).Error
	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "产品删除失败",
			Error: err.Error(),
		}
	}

	return serializer.SuccessResponse()

}

// ListSpotService 全部景点信息结构体
type ListProductService struct {
	Start int `form:"start" json:"start"`
	Limit int `form:"limit" json:"limit"`
}

// List 景点列表
func (service *ListProductService) ListProduct(c *gin.Context) serializer.Response {
	var infos []model.ProductInfo
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
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}

	return serializer.BuildProductResponse(infos)
}

// GetProductService 获取产品的服务
type GetProductService struct {
}

// GetProductByID 根据景点id获取产品信息
func (service *GetProductService) GetProductByID(c *gin.Context, id string) serializer.Response {
	var spot model.SpotInfo
	err := model.DB.First(&spot, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "景点不存在",
			Error: err.Error(),
		}
	}
	var products []model.ProductInfo
	err = model.DB.Where("s_id=?", spot.ID).Find(&products).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "产品不存在",
			Error: err.Error(),
		}
	}

	return serializer.BuildProductResponse(products)
}

// CollectSpotService 收藏产品信息结构体
type CollectProductService struct {
}

func (service *CollectProductService) CollectProduct(c *gin.Context, pid string) serializer.Response {
	//产品是否存在
	var product model.ProductInfo
	err := model.DB.First(&product, pid).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "产品不存在",
			Error: err.Error(),
		}
	}
	if c_user, _ := c.Get("user"); c_user != nil {
		if u, ok := c_user.(*model.User); ok {
			skey := "product_set_" + u.UserName
			zkey := "product_zset_" + u.UserName
			score := time.Now().Unix()
			//是否已经收藏
			isExist := cache.RedisClient.SIsMember(skey, pid).Val()
			if !isExist {
				cache.RedisClient.SAdd(skey, pid)
				cache.RedisClient.ZAdd(zkey, redis.Z{
					Score:  float64(score),
					Member: pid,
				})
			} else {
				return serializer.Response{
					Code: 400,
					Msg:  "产品已收藏",
				}
			}

		}
	}
	return serializer.SuccessResponse()
}

// CollectSpotService 收藏产品信息结构体
type LikeProductService struct {
}

func (service *LikeProductService) LikeProduct(c *gin.Context, pid string) serializer.Response {
	//产品是否存在
	var product model.ProductInfo
	err := model.DB.First(&product, pid).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "产品不存在",
			Error: err.Error(),
		}
	}
	if c_user, _ := c.Get("user"); c_user != nil {
		if u, ok := c_user.(*model.User); ok {
			skey := "product_like_" + u.UserName
			key := "product_like_" + pid
			//是否已经点赞
			isExist := cache.RedisClient.SIsMember(skey, pid).Val()
			if !isExist {
				cache.RedisClient.SAdd(skey, pid)
				cache.RedisClient.IncrBy(key, 1)
			} else {
				return serializer.Response{
					Code: 400,
					Msg:  "产品已点赞",
				}
			}

		}
	}
	return serializer.SuccessResponse()
}

func (service *LikeProductService) DisLikeProduct(c *gin.Context, pid string) serializer.Response {
	//产品是否存在
	var product model.ProductInfo
	err := model.DB.First(&product, pid).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "产品不存在",
			Error: err.Error(),
		}
	}
	if c_user, _ := c.Get("user"); c_user != nil {
		if u, ok := c_user.(*model.User); ok {
			skey := "product_like_" + u.UserName
			key := "product_like_" + pid
			//是否已经点赞
			isExist := cache.RedisClient.SIsMember(skey, pid).Val()
			if isExist {
				cache.RedisClient.SRem(skey, pid)
				cache.RedisClient.IncrBy(key, -1)
			} else {
				return serializer.Response{
					Code: 400,
					Msg:  "产品尚未点赞",
				}
			}

		}
	}
	return serializer.SuccessResponse()
}
