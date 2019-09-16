package service

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
)

type GetMsgService struct {

}

func (GetMsgService) Get(id string) *serializer.Response {
	var msg model.Msg
	// 查询
	if err := model.DB.First(&msg, id).Error; err != nil {
		return &serializer.Response{
			Status:    40001,
			Msg:       "没有这条信息",
		}
	}

	return &serializer.Response{
		Status:    0,
		Data:      serializer.OneMsgResponse(msg),
	}
}