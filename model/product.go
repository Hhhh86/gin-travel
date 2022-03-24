package model

import (
	"gorm.io/gorm"
	"time"
)

// ProductInfo 景点信息
type ProductInfo struct {
	gorm.Model
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
