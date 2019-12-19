package service

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

type UploadCloudItem struct {
}

func SaveFile(filename string, file multipart.File) error {
	// 打开文件句柄
	openFile, err := os.OpenFile(fmt.Sprintf("upload/%s", filename), os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer openFile.Close()
	if err != nil {
		return err
	}

	// 创建缓冲区
	buf := make([]byte, 4096)
	for {
		_, err := file.Read(buf)
		// 读取完毕就退出循环
		if err == io.EOF {
			break
		}
		_, err = openFile.Write(buf)
	}

	return nil
}

func (UploadCloudItem) Upload(file *multipart.File, header *multipart.FileHeader) *serializer.Response {
	filename := header.Filename
	size := int(header.Size / 1024)
	fileCloud := model.FileCloud{
		Size:     size,
		FileName: filename,
	}
	count := 0
	model.DB.Model(&model.FileCloud{}).Where("file_name = ?", filename).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: serializer.UserInputError,
			Msg:  "该条纪录已经存在",
		}
	}

	if err := model.DB.Create(&fileCloud).Error; err != nil {
		return &serializer.Response{
			Code: serializer.DatabaseWriteError,
			Msg:  "数据库写入错误.",
		}
	}

	// 保存文件到本地 传入filename, file
	err := SaveFile(filename, *file)
	if err != nil {
		return &serializer.Response{
			Code:  serializer.ServerPanicError,
			Msg:   "保存上传文件失败.",
			Error: fmt.Sprintf("%v", err),
		}
	}

	return &serializer.Response{
		Data: serializer.FileCloudSerializeResponse(fileCloud),
	}
}
