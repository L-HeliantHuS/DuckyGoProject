package service

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
)

type GetCloudListService struct {
}

func (GetCloudListService) Gets() *serializer.Response {
	var files []model.FileCloud
	var count int
	if err := model.DB.Find(&files).Error; err != nil {
		return &serializer.Response{
			Code: serializer.DatabaseReadError,
			Msg:  "DataBase Read Error!",
		}
	}

	if err := model.DB.Table("file_clouds").Count(&count).Error; err != nil {
		return &serializer.Response{
			Code:      serializer.DatabaseReadError,
			Msg:       "Database Read Count Error!",
		}
	}

	return &serializer.Response{
		Data: serializer.FileCloudSerializerAllResponse(files, count),
	}
}
