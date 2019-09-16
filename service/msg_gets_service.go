package service

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
)

type GetMegsService struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

func (service *GetMegsService) Gets() *serializer.Response {
	var msgs []model.Msg
	count := 0
	// 没有传递参数就返回20条数据
	if service.Limit == 0 {
		service.Limit = 20
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Offset).Find(&msgs).Error; err != nil {
		return &serializer.Response{
			Status: 50001,
			Msg:    "数据库操作异常",
		}
	}

	if err := model.DB.Table("msgs").Count(&count).Error; err != nil {
		return &serializer.Response{
			Status: 50001,
			Msg:    "数据库操作异常",
		}
	}


	return &serializer.Response{
		Data: serializer.MsgsResponse(msgs, count),
	}
}
