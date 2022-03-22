package server

import (
	"os"
	"singo/api"
	"singo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())
	r.Use(middleware.CurrentBusiness())

	//127.0.0.1/ping
	r.GET("ping", api.Ping)
	//todo 全部旅游景点页
	r.GET("/all/spot", api.AllSpot)
	//todo 全部旅游产品页
	r.GET("/all/products", api.UserMe)
	//todo 旅游攻略页
	r.GET("/detail/list", api.UserMe)
	// 路由
	v1 := r.Group("/api/user")
	{
		// 用户注册
		v1.POST("/register", api.UserRegister)

		// 用户登录
		v1.POST("/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("", middleware.AuthRequired(true))
		{
			// User Routing
			auth.GET("/me", api.UserMe)
			auth.DELETE("/logout", api.UserLogout)
			//todo 个人资料 crud
			auth.POST("/create/info", api.UserLogin)
			auth.DELETE("/delete/info", api.UserLogin)
			auth.POST("/retrieve/info", api.UserLogin)
			auth.PUT("/update/info", api.UserLogin)

			//todo 推荐旅游产品页
			auth.GET("/recommend/list", api.UserMe)
			//todo 报名
			auth.POST("/apply", api.UserMe)
			//todo 查看订单
			auth.GET("/retrieve/order", api.UserMe)
			//todo 取消订单
			auth.DELETE("/cancel/order", api.UserMe)
			//todo 评价订单
			auth.POST("/evaluate/order", api.UserMe)
			//todo 发布攻略
			auth.POST("/publish/strategy", api.UserMe)
			//todo 评论
			auth.POST("/evaluate/strategy", api.UserMe)
			//todo 点赞
			auth.GET("/like/strategy", api.UserMe)
			//todo 取消点赞
			auth.GET("/dislike/strategy", api.UserMe)
			//todo 抽奖
			auth.POST("/draw/beat", api.UserMe)
			//todo 获奖记录
			auth.POST("/draw/rewards", api.UserMe)
		}
	}

	v2 := r.Group("/api/business")
	{
		// 商家注册
		v2.POST("/register", api.BusinessRegister)

		// 商家登录
		v2.POST("/login", api.BusinessLogin)

		// 需要登录保护的
		auth := v2.Group("", middleware.AuthBusinessRequired(false))
		{
			// business Routing
			auth.GET("/me", api.UserMe)            //todo
			auth.DELETE("/logout", api.UserLogout) //todo
			//todo 旅游景点资料 crud
			auth.POST("/create/spot", api.CreateSpot)
			auth.DELETE("/delete/spot/:id", api.DeleteSpot)
			auth.PUT("/update/spot/:id", api.UpdateSpot)
			//todo 旅游产品资料 crud
			auth.POST("/create/products", api.UserLogin)
			auth.DELETE("/delete/products", api.UserLogin)
			auth.POST("/retrieve/products", api.UserLogin)
			auth.PUT("/update/products", api.UserLogin)
			//todo 查看订单信息
			auth.GET("/retrieve/order", api.UserLogin)
			//todo 处理订单(接受/拒绝)
			auth.POST("/handle/order", api.UserLogin)
			//todo 查看订单评价
			auth.POST("/retrieve/evaluation", api.UserLogin)
		}
	}
	return r
}
