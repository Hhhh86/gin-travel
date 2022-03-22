package serializer

import "singo/model"

// Business 用户序列化器
type Business struct {
	ID          uint   `json:"id"`
	CompanyName string `json:"company_name"`
	Status      string `json:"status"`
	Avatar      string `json:"avatar"`
	CreatedAt   int64  `json:"created_at"`
}

// BuildUser 序列化用户
func BuildBusiness(business model.Business) Business {
	return Business{
		ID:          business.ID,
		CompanyName: business.CompanyName,
		Status:      business.Status,
		Avatar:      business.Avatar,
		CreatedAt:   business.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildBusinessResponse(business model.Business) Response {
	return Response{
		Data: BuildBusiness(business),
	}
}
