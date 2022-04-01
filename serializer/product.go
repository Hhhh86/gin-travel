package serializer

import (
	"singo/model"
	"time"
)

// ProductInfo 产品序列化器
type ProductInfo struct {
	SID         int       //景点id
	BID         int       //公司id
	Number      int       //人数
	Price       int       //套餐价格
	CompanyName string    //公司名称
	Name        string    //套餐名称
	Describe    string    //描述
	StartTime   time.Time //出发时间
	EndTime     time.Time //结束时间
}

// BuildProduct 序列化景点
func BuildProduct(product model.ProductInfo) ProductInfo {
	return ProductInfo{
		SID:         product.SID,
		BID:         product.BID,
		Number:      product.Number,
		Price:       product.Price,
		CompanyName: product.CompanyName,
		Name:        product.Name,
		Describe:    product.Describe,
		StartTime:   product.StartTime,
		EndTime:     product.EndTime,
	}
}

// BuildUserResponse 序列化用户响应
func BuildProductResponse(infos []model.ProductInfo) Response {
	return Response{
		Data: BuildProductList(infos),
	}
}

// BuildSpotList 序列化景点列表
func BuildProductList(items []model.ProductInfo) (infos []ProductInfo) {
	for _, item := range items {
		info := BuildProduct(item)
		infos = append(infos, info)
	}
	return infos
}
