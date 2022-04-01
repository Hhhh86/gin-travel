package model

import "gorm.io/gorm"

//景点攻略
type StrategyInfo struct {
	gorm.Model
	SID      int    //景点id
	Like     int    //点赞数
	Content  string //攻略内容
	UserName string //用户名称
	Photos   string //图片
}

//攻略评论
type StrategyCommentInfo struct {
	gorm.Model
	STID     int    //攻略id
	Like     int    //点赞数
	UserName string //用户名称
	Content  string //内容
}
