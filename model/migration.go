package model

//执行数据迁移

func migration() {
	// 自动迁移模式
	_ = DB.AutoMigrate(&User{})        //用户登陆信息
	_ = DB.AutoMigrate(&Business{})    //商家登陆信息
	_ = DB.AutoMigrate(&UserInfo{})    //用户个人信息
	_ = DB.AutoMigrate(&SpotInfo{})    //景点信息
	_ = DB.AutoMigrate(&ProductInfo{}) //产品信息

}
