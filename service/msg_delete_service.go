package service

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
)

type DeleteMsgService struct {
}

func (DeleteMsgService) Delete(id string, userid uint) *serializer.Response {
	var msg model.Msg

	// 从数据库查询出来数据
	if err := model.DB.First(&msg, id).Error; err != nil {
		return &serializer.Response{
			Status: 40001,
			Msg:    "没有这条数据",
		}
	}

	// 检测是否是当前用户创建的留言
	if msg.UserID != userid {
		return &serializer.Response{
			Status: 40003,
			Msg:    "这条留言不是你的！",
		}
	}

	if err := model.DB.Delete(&msg).Error; err != nil {
		return &serializer.Response{
			Status: 50003,
			Msg:    "数据库删除错误",
		}
	}

	return &serializer.Response{
		Data:      serializer.MsgResponse(msg),
		Msg:       "删除成功",
	}
}
