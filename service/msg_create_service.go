package service

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
)

type CreateMsgService struct {
	Message string `form:"msg" json:"msg" binding:"required"`
	Subject string `form:"til" json:"til" binding:"required"`
}

func (service *CreateMsgService) Create(userID uint) *serializer.Response {

	msg := model.Msg{
		UserID:  userID,
		Message: service.Message,
		Subject: service.Subject,
	}

	if err := model.DB.Create(&msg).Error; err != nil {
		return &serializer.Response{
			Status: 50001,
			Msg:    "保存失败",
		}
	}

	return &serializer.Response{
		Data: serializer.OneMsgResponse(msg),
		Msg:  "",
	}
}
