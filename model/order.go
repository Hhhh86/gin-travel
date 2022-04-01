package model

import (
	"gorm.io/gorm"
)

// OrderInfo 订单信息
type OrderInfo struct {
	gorm.Model
	PID      int    //产品id
	Status   int    //状态
	Number   int    //人数
	UserName string //用户名称
}

//订单评论
type OrderComment struct {
	gorm.Model
	OID      int    //订单id
	Content  string //评论内容
	UserName string //用户名称
}
