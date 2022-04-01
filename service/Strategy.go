package service

import (
	"github.com/gin-gonic/gin"
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
