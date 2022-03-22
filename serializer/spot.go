package serializer

import (
	"singo/model"
	"time"
)

// SpotInfo 景点序列化器
type SpotInfo struct {
	ID           uint  `json:"id"`
	AgeScope     int   //0Children、1Youth、2Midlife，3Aged
	PlaceType    int   //1自然生态类、2历史文化类、3现代游乐类、4产业融合类、5其他类。
	Price        int   //门票价格
	CreatedAt    time.Time `json:"created_at"`
	Name         string
	Address      string
	Introduction string
	Tips         string
	OpenTime     string
}

// BuildUser 序列化用户
func BuildSpot(spot model.SpotInfo) SpotInfo {
	return SpotInfo{
		ID:           spot.ID,
		Name:         spot.Name,
		Address:      spot.Address,
		AgeScope:     spot.AgeScope,
		Introduction: spot.Introduction,
		Tips:         spot.Tips,
		OpenTime:     spot.OpenTime,
		PlaceType:    spot.PlaceType,
		Price:        spot.Price,
	}
}

// BuildUserResponse 序列化用户响应
func BuildSpotResponse(infos []model.SpotInfo) Response {
	return Response{
		Data: BuildSpotList(infos),
	}
}


// BuildSpotList 序列化景点列表
func BuildSpotList(items []model.SpotInfo) (infos []SpotInfo) {
	for _, item := range items {
		info := BuildSpot(item)
		infos = append(infos, info)
	}
	return infos
}