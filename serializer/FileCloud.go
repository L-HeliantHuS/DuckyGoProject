package serializer

import "DuckyGo/model"

// FileCloudSerializer 网盘文件序列化器
type FileCloudSerializer struct {
	ID       uint   `json:"id"`
	FileName string `json:"filename"`
	Size     int    `json:"size"`
	CreateAt int64  `json:"create_at"`
}

// FileCloudSerializerAll 网盘文件多个数据序列化器
type FileCloudSerializerAll struct {
	Results []FileCloudSerializer `json:"list"`
	Count   int                   `json:"count"`
}

// FileCloudSerializeResponse 网盘文件序列化响应
func FileCloudSerializeResponse(db model.FileCloud) FileCloudSerializer {
	return FileCloudSerializer{
		ID:       db.ID,
		FileName: "/upload/" + db.FileName,
		Size:     db.Size,
		CreateAt: db.CreatedAt.Unix(),
	}
}

// BaseAllResponse 多个数据序列化响应
func FileCloudSerializerAllResponse(db []model.FileCloud, count int) FileCloudSerializerAll {
	var result []FileCloudSerializer
	for _, i := range db {
		result = append(result, FileCloudSerializeResponse(i))
	}
	return FileCloudSerializerAll{
		Results: result,
		Count:   count,
	}
}
