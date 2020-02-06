package serializer

import "DuckyGo/model"

type ImageSerializer struct {
	ID         uint   `json:"id"`
	ImageTitle string `json:"title"`
	ImageURL   string `json:"url"`
}

type ImageAllSerializer struct {
	List  []ImageSerializer `json:"list"`
	Count int               `json:"count"`
}

// ImageResponse 单个图片响应
func ImageResponse(db model.Image) ImageSerializer {
	return ImageSerializer{
		ID:         db.ID,
		ImageTitle: db.ImageTitle,
		ImageURL:   db.SignURL(),
	}
}

// ImageAllResponse 多个图片响应
func ImageAllResponse(db []model.Image, count int) ImageAllSerializer {
	var result []ImageSerializer
	for _, i := range db {
		result = append(result, ImageResponse(i))
	}
	return ImageAllSerializer{
		List:  result,
		Count: count,
	}
}
