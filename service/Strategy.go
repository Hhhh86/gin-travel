package service

import (
	"github.com/gin-gonic/gin"
	"singo/cache"
	"singo/model"
	"singo/serializer"
)

// CreateStrategyService 新建攻略信息结构体
type CreateStrategyService struct {
	SID     int    `json:"s_id"`    //景点id
	Content string `json:"content"` //攻略内容
	Photos  string `json:"photos"`  //图片
}

// CreateStrategy 新建攻略信息
func (service *CreateStrategyService) CreateStrategy(c *gin.Context) serializer.Response {
	var spot model.SpotInfo
	err := model.DB.First(&spot, service.SID).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "景点不存在",
			Error: err.Error(),
		}
	}
	if c_user, _ := c.Get("user"); c_user != nil {
		if u, ok := c_user.(*model.User); ok {
			// 新建攻略
			info := model.StrategyInfo{
				SID:      service.SID,
				Content:  service.Content,
				UserName: u.UserName,
				Photos:   service.Photos,
			}
			if err := model.DB.Create(&info).Error; err != nil {
				return serializer.ParamErr("新建攻略失败", err)
			}

		}
	}
	return serializer.SuccessResponse()
}

// DeleteStrategyService 删除攻略信息结构体
type DeleteStrategyService struct {
}

// DeleteStrategy 删除攻略信息
func (service *DeleteStrategyService) DeleteStrategy(c *gin.Context, id string) serializer.Response {
	if c_user, _ := c.Get("user"); c_user != nil {
		if u, ok := c_user.(*model.User); ok {
			var strategy model.StrategyInfo
			//查询是否有该攻略
			err := model.DB.Where("id = ? and user_name = ?", id, u.UserName).First(&strategy).Error
			if err != nil {
				return serializer.Response{
					Code:  404,
					Msg:   "攻略不存在",
					Error: err.Error(),
				}
			}

			err = model.DB.Delete(&strategy).Error
			if err != nil {
				return serializer.Response{
					Code:  500,
					Msg:   "删除攻略错误",
					Error: err.Error(),
				}
			}

		}
	}
	return serializer.SuccessResponse()
}

// CreateStrategyCommentService 评论攻略信息结构体
type CreateStrategyCommentService struct {
	STID    int    `json:"st_id"`   //攻略id
	Content string `json:"content"` //内容
}

// CreateComment 评论攻略信息
func (service *CreateStrategyCommentService) CreateComment(c *gin.Context) serializer.Response {
	var strategy model.StrategyInfo
	err := model.DB.First(&strategy, service.STID).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "攻略不存在",
			Error: err.Error(),
		}
	}
	if c_user, _ := c.Get("user"); c_user != nil {
		if u, ok := c_user.(*model.User); ok {
			// 新建攻略
			info := model.StrategyCommentInfo{
				STID:     service.STID,
				Content:  service.Content,
				UserName: u.UserName,
			}
			if err := model.DB.Create(&info).Error; err != nil {
				return serializer.ParamErr("新建攻略评论失败", err)
			}
		}
	}
	return serializer.SuccessResponse()
}

// DeleteStrategyService 删除攻略信息结构体
type DeleteStrategyCommentService struct {
}

// DeleteComment 删除攻略信息
func (service *DeleteStrategyCommentService) DeleteComment(c *gin.Context, id string) serializer.Response {
	if c_user, _ := c.Get("user"); c_user != nil {
		if u, ok := c_user.(*model.User); ok {
			var comment model.StrategyCommentInfo
			//查询是否有该攻略
			err := model.DB.Where("id = ? and user_name = ?", id, u.UserName).First(&comment).Error
			if err != nil {
				return serializer.Response{
					Code:  404,
					Msg:   "评论不存在",
					Error: err.Error(),
				}
			}

			err = model.DB.Delete(&comment).Error
			if err != nil {
				return serializer.Response{
					Code:  500,
					Msg:   "删除评论错误",
					Error: err.Error(),
				}
			}

		}
	}
	return serializer.SuccessResponse()
}

// LikeStrategyService 收藏产品信息结构体
type LikeStrategyService struct {
}

func (service *LikeStrategyService) LikeStrategy(c *gin.Context, id string) serializer.Response {
	//攻略是否存在
	var strategy model.StrategyInfo
	err := model.DB.First(&strategy, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "攻略不存在",
			Error: err.Error(),
		}
	}
	//string 用户积分
	ukey := "score_" + strategy.UserName
	if c_user, _ := c.Get("user"); c_user != nil {
		if u, ok := c_user.(*model.User); ok {
			//set 用户是否有点赞
			skey := "strategy_like_" + u.UserName
			//string 攻略的点赞
			key := "strategy_like_" + id
			//是否已经点赞
			isExist := cache.RedisClient.SIsMember(skey, id).Val()
			if !isExist {
				cache.RedisClient.SAdd(skey, id)
				cache.RedisClient.IncrBy(key, 1)
				//增加积分
				cache.RedisClient.IncrBy(ukey, 1)
			} else {
				return serializer.Response{
					Code: 400,
					Msg:  "攻略已点赞",
				}
			}

		}
	}
	return serializer.SuccessResponse()
}

func (service *LikeStrategyService) DisLikeStrategy(c *gin.Context, id string) serializer.Response {
	var strategy model.StrategyInfo
	err := model.DB.First(&strategy, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "攻略不存在",
			Error: err.Error(),
		}
	}
	//string 用户积分
	ukey := "score_" + strategy.UserName
	if c_user, _ := c.Get("user"); c_user != nil {
		if u, ok := c_user.(*model.User); ok {
			skey := "strategy_like_" + u.UserName
			key := "strategy_like_" + id
			//是否已经点赞
			isExist := cache.RedisClient.SIsMember(skey, id).Val()
			if isExist {
				cache.RedisClient.SRem(skey, id)
				cache.RedisClient.IncrBy(key, -1)
				cache.RedisClient.IncrBy(ukey, -1)
			} else {
				return serializer.Response{
					Code: 400,
					Msg:  "攻略尚未点赞",
				}
			}

		}
	}
	return serializer.SuccessResponse()
}
