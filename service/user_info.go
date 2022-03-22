package service

// UserInfoService 管理用户登录的服务
type UserInfoService struct {
	AgeScope    string `form:"age_scope" json:"age_scope" binding:"required,min=5,max=30"`
	ExpectPrice string `form:"expect_price" json:"expect_price" binding:"required,min=8,max=40"`
	Sex         string `form:"sex" json:"sex" binding:"required,min=8,max=40"`
}
