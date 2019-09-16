package service

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
)

type GetsHouseService struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

func (service *GetsHouseService) Gets() *serializer.Response {
	var lianjia []model.Lianjia

	count := 0
	// 没有传递参数就返回20条数据
	if service.Limit == 0 {
		service.Limit = 20
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Offset).Find(&lianjia).Error; err != nil {
		return &serializer.Response{
			Status: 50001,
			Msg:    "数据库操作异常",
		}
	}

	if err := model.DB.Table("lianjia").Count(&count).Error; err != nil {
		return &serializer.Response{
			Status: 50001,
			Msg:    "数据库操作异常",
		}
	}

	return &serializer.Response{
		Data: serializer.LianjiaDatasResponse(lianjia, count),
	}
}
