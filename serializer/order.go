package serializer

import (
	"singo/model"
	"time"
)

type OrderAllInfo struct {
	ID       int    `json:"id"`
	PID      int    `json:"p_id"`      //产品id
	UserName string `json:"user_name"` //用户名称
	Number   int    `json:"number"`    //人数
	Status   int    `json:"status"`    //人数
}

// BuildProduct 序列化景点
func BuildAllOrder(order model.OrderInfo) OrderAllInfo {
	return OrderAllInfo{
		ID:       int(order.ID),
		PID:      order.PID,
		UserName: order.UserName,
		Number:   order.Number,
		Status:   order.Status,
	}
}

// BuildUserResponse 序列化用户响应
func BuildAllOrderResponse(infos []model.OrderInfo) Response {
	return Response{
		Data: BuildOrderList(infos),
	}
}

// BuildSpotList 序列化景点列表
func BuildOrderList(items []model.OrderInfo) (infos []OrderAllInfo) {
	for _, item := range items {
		info := BuildAllOrder(item)
		infos = append(infos, info)
	}
	return infos
}

type OrderOneInfo struct {
	ID          int       `json:"id"`
	PID         int       `json:"p_id"`         //产品id
	UserName    string    `json:"user_name"`    //用户名称
	Number      int       `json:"number"`       //人数
	Status      int       `json:"status"`       //状态
	SID         int       `json:"s_id"`         //景点id
	BID         int       `json:"b_id"`         //公司id
	Price       int       `json:"price"`        //套餐价格
	CompanyName string    `json:"company_name"` //公司名称
	Name        string    `json:"name"`         //套餐名称
	Describe    string    `json:"describe"`     //描述
	StartTime   time.Time `json:"start_time"`   //出发时间
	EndTime     time.Time `json:"end_time"`     //结束时间
}

// BuildProduct 序列化景点
func BuildOneOrder(order model.OrderInfo, product model.ProductInfo) OrderOneInfo {
	return OrderOneInfo{
		ID:          int(order.ID),
		PID:         order.PID,
		UserName:    order.UserName,
		Number:      order.Number,
		Status:      order.Status,
		SID:         product.SID,
		BID:         product.BID,
		Price:       product.Price,
		CompanyName: product.CompanyName,
		Name:        product.Name,
		Describe:    product.Describe,
		StartTime:   product.StartTime,
		EndTime:     product.EndTime,
	}
}

// BuildOneOrderResponse 序列化用户响应
func BuildOneOrderResponse(order model.OrderInfo, product model.ProductInfo) Response {
	return Response{
		Data: BuildOneOrder(order, product),
	}
}
