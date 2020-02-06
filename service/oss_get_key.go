package service

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	uuid "github.com/satori/go.uuid"
	"os"
	"strings"
)

type GetOSSKeyService struct {
	Filename string `form:"filename" json:"filename" binding:"required"`
}

func SaveToDB(title string, filename string) error {
	image := model.Image{
		ImageTitle: title,
		ImageSign:  filename,
	}
	if err := model.DB.Create(&image).Error; err != nil {
		return err
	}

	return nil
}

func (service *GetOSSKeyService) POST() *serializer.Response {
	client, err := oss.New(os.Getenv("ALIYUN_OSS_END_POINT"), os.Getenv("ALIYUN_OSS_ACCESS_KEY"), os.Getenv("ALIYUN_OSS_SECRET_KEY"))
	if err != nil {
		return &serializer.Response{
			Code:  serializer.ServerPanicError,
			Msg:   "OSS初始化失败.",
			Error: err.Error(),
		}
	}

	// 获取Bucket目录
	bucket, err := client.Bucket(os.Getenv("ALIYUN_OSS_BUCKET"))
	if err != nil {
		return &serializer.Response{
			Code:  serializer.ServerPanicError,
			Msg:   "Bucket可能不存在.",
			Error: err.Error(),
		}
	}

	// 设置参数
	options := []oss.Option{
		oss.ContentType("image/png"),
	}

	// 上传路径
	uploadAddr := "images/" + uuid.Must(uuid.NewV4(), nil).String() + ".png"

	// 签名PUT上传URL
	signPutURL, err := bucket.SignURL(uploadAddr, oss.HTTPPut, 300, options...)
	if err != nil {
		return &serializer.Response{
			Code:  serializer.ServerPanicError,
			Msg:   "签名上传URL失败",
			Error: err.Error(),
		}
	}

	// 签名GET获取URL
	signGetURL, err := bucket.SignURL(uploadAddr, oss.HTTPGet, 300)
	if err != nil {
		return &serializer.Response{
			Code:  serializer.ServerPanicError,
			Msg:   "签名获取URL失败",
			Error: err.Error(),
		}
	}

	// 把保存的图片保存到数据库
	title := strings.Split(service.Filename, ".")[0]
	err = SaveToDB(title, uploadAddr)
	if err != nil {
		return &serializer.Response{
			Code:  serializer.DatabaseWriteError,
			Msg:   "数据库写入失败.",
			Error: err.Error(),
		}
	}

	return &serializer.Response{
		Code: 0,
		Data: map[string]string{
			"upload": uploadAddr,
			"put":    signPutURL,
			"get":    signGetURL,
		},
	}
}
