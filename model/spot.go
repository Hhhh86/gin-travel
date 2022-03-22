package model

import (
	"gorm.io/gorm"
)

// SpotInfo 景点信息
type SpotInfo struct {
	gorm.Model
	AgeScope     int   //0Children、1Youth、2Midlife，3Aged
	PlaceType    int   //1自然生态类、2历史文化类、3现代游乐类、4产业融合类、5其他类。
	Price        int   //门票价格
	Name         string
	Address      string
	Introduction string
	Tips         string
	OpenTime     string
}
