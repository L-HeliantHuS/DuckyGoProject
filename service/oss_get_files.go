package service

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
)

type GetOSSFileListService struct {
}

func (GetOSSFileListService) Get() *serializer.Response {
	var images []model.Image
	var count int

	// 查询数据
	if err := model.DB.Where(model.Image{}).Find(&images).Error; err != nil {
		return &serializer.Response{
			Code:  serializer.DatabaseReadError,
			Msg:   "数据库读取失败",
			Error: err.Error(),
		}
	}

	// 读取Count
	if err := model.DB.Table("images").Count(&count).Error; err != nil {
		return &serializer.Response{
			Code:  serializer.DatabaseReadError,
			Msg:   "数据库读取失败",
			Error: err.Error(),
		}
	}

	return &serializer.Response{
		Data: serializer.ImageAllResponse(images, count),
	}
}
