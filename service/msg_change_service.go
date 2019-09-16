package service

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
)

type ChangeMsgService struct {
	Message string `form:"msg" json:"msg"`
	Subject string `form:"til" json:"til"`
}

func (service *ChangeMsgService) Change(id string, userid uint) *serializer.Response {
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

	// 修改这条信息
	if err := model.DB.Model(&msg).Updates(model.Msg{
		Message: service.Message,
		Subject: service.Subject,
	}).Error; err != nil {
		return &serializer.Response{
			Status: 50002,
			Msg:    "数据库操作异常",
		}
	}

	return &serializer.Response{
		Data: serializer.MsgResponse(msg),
		Msg:  "更新成功！",
	}
}
