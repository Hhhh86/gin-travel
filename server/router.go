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
	r.Use(middleware.CurrentUserAndBusiness())

	//127.0.0.1/ping
	r.GET("ping", api.Ping)
	//todo 全部旅游景点页
	r.GET("/spot/all", api.AllSpot) //done
	//todo 全部旅游产品页
	r.GET("/products/all", api.AllProduct) //done
	//todo 根据sid查询对应的产品
	r.GET("/products/retrieve/:id", api.GetProductByID) //done
	//todo 旅游攻略页
	r.GET("/list/detail", api.UserMe)
	// 路由
	v1 := r.Group("/api/user")
	{
		//done 用户注册
		v1.POST("/register", api.UserRegister)

		//done 用户登录
		v1.POST("/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("", middleware.AuthRequired(true))
		{
			// User Routing
			auth.GET("/me", api.UserMe)
			auth.DELETE("/logout", api.UserLogout)
			//todo 个人资料 crud
			auth.POST("/info/create", api.CreateInfo)    //done
			auth.DELETE("/info/delete", api.DeleteInfo)  //done
			auth.GET("/info/retrieve", api.RetrieveInfo) //done
			auth.PUT("/info/update", api.UpdateInfo)     //done
			//todo 推荐用户
			//todo 添加好友
			//todo 用户私信
			//todo 推荐旅游产品页
			auth.GET("/list/recommend", api.UserMe)
			//todo 收藏景点
			auth.GET("/spot/collect/:sid", api.CollectSpot) //done
			//todo 收藏产品
			auth.GET("/product/collect/:pid", api.CollectProduct) //done
			//todo 点赞产品
			auth.GET("/product/like/:pid", api.LikeProduct) //done
			//todo 取消点赞产品
			auth.GET("/product/dislike/:pid", api.DisLikeProduct) //done
			//todo 查看收藏
			auth.GET("/collect/retrieve/:tag", api.UserMe)
			//todo 报名
			auth.POST("/order/create", api.CreateOrder) //done
			//todo 查看全部订单
			auth.GET("/order/all/retrieve", api.RetrieveAllOrder) //done
			//todo 查看单一订单
			auth.GET("/order/one/retrieve/:id", api.RetrieveOneOrder) //done
			//todo 取消订单
			auth.POST("/order/cancel/:id", api.CancelOrder) //done
			//todo 删除订单
			auth.DELETE("/order/delete/:id", api.DeleteOrder) //done
			//todo 评价订单
			auth.POST("/order/evaluate", api.EvaluateOrder) //done
			//todo 发布景点攻略
			auth.POST("/strategy/create", api.CreateStrategy) //done
			//todo 删除景点攻略
			auth.DELETE("/strategy/delete/:sid", api.DeleteStrategy) //done
			//todo 评论攻略
			auth.POST("/strategy/comment/create", api.CreateStrategyComment) //done
			//todo 删除攻略评论
			auth.DELETE("/strategy/comment/delete/:stid", api.DeleteStrategyComment) //done
			//todo 点赞攻略
			auth.GET("/strategy/like/:eid", api.UserMe)
			//todo 取消点赞攻略
			auth.GET("/strategy/dislike/:eid", api.UserMe)
			//todo 抽奖
			auth.POST("/draw/beat", api.UserMe)
			//todo 获奖记录
			auth.POST("/draw/rewards", api.UserMe)
		}
	}

	v2 := r.Group("/api/business")
	{
		// 商家注册
		v2.POST("/register", api.BusinessRegister) //done

		// 商家登录
		v2.POST("/login", api.BusinessLogin) //done

		// 需要登录保护的
		auth := v2.Group("", middleware.AuthBusinessRequired(false))
		{
			// business Routing
			auth.GET("/me", api.BusinessMe)            //done
			auth.DELETE("/logout", api.BusinessLogout) //done
			//todo 旅游景点资料 crud
			auth.POST("/spot/create", api.CreateSpot)       //done
			auth.DELETE("/spot/delete/:id", api.DeleteSpot) //done
			auth.PUT("/spot/update/:id", api.UpdateSpot)    //done
			//todo 旅游产品资料 crud
			auth.POST("/product/create", api.CreateProduct)       //done
			auth.DELETE("/product/delete/:id", api.DeleteProduct) //done
			auth.PUT("/product/update", api.UpdateProduct)        //done
			//todo 查看订单信息
			auth.GET("/order/retrieve", api.UserLogin)
			//todo 处理订单(接受/拒绝)
			auth.POST("/order/handle", api.UserLogin)
			//todo 查看订单评价
			auth.GET("/evaluate/retrieve", api.UserLogin)
		}
	}
	return r
}
