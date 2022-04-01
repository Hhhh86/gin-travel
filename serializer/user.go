package serializer

import (
	"singo/model"
)

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	NickName  string `json:"nick_name"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		NickName:  user.NickName,
		Status:    user.Status,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}

type UserInfo struct {
	Sex            int    `json:"sex"`              //1女 2男
	PlaceType      int    `json:"place_type"`       //1自然生态类、2历史文化类、3现代游乐类、4产业融合类、5其他类。
	AgeScope       int    `json:"age_scope"`        //0Children、1Youth、2Midlife，3Aged
	ExpectMinPrice int    `json:"expect_min_price"` //期望最小价格
	ExpectMaxPrice int    `json:"expect_max_price"` //期望最大价格
	UserName       string `json:"user_name"`
	NickName       string `json:"nick_name"`
	Avatar         string `json:"avatar"`
}

// BuildUserInfo 序列化用户信息
func BuildUserInfo(info model.UserInfo) UserInfo {
	return UserInfo{
		Sex:            info.Sex,
		PlaceType:      info.PlaceType,
		AgeScope:       info.AgeScope,
		ExpectMinPrice: info.ExpectMinPrice,
		ExpectMaxPrice: info.ExpectMaxPrice,
		UserName:       info.UserName,
		NickName:       info.NickName,
		Avatar:         info.Avatar,
	}
}

// BuildUserInfoResponse 序列化用户响应
func BuildUserInfoResponse(info model.UserInfo) Response {
	return Response{
		Data: BuildUserInfo(info),
	}
}
