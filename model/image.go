package model

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
	"os"
)

type Image struct {
	gorm.Model
	ImageTitle string
	ImageSign  string
}

func (model *Image) SignURL() string {
	client, _ := oss.New(os.Getenv("ALIYUN_OSS_END_POINT"), os.Getenv("ALIYUN_OSS_ACCESS_KEY"), os.Getenv("ALIYUN_OSS_SECRET_KEY"))
	bucket, _ := client.Bucket(os.Getenv("ALIYUN_OSS_BUCKET"))
	resultURL, _ := bucket.SignURL(model.ImageSign, oss.HTTPGet, 300)
	return resultURL
}
