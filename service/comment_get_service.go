package service

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
	"strconv"
)

type GetCommentService struct {
}

const PageSize int = 5

// TypeChange 类型转换
func TypeChange(oid string, pn string) (int, int, error) {
	oidInt, err := strconv.Atoi(oid)
	if err != nil {
		return 0, 0, err
	}

	pnInt, err := strconv.Atoi(pn)
	if err != nil {
		return 0, 0, err
	}

	return oidInt, pnInt, nil

}

// Get 根据传递过来的oid获取下面所有的评论
func (service *GetCommentService) Get(oidStr string, pnStr string) *serializer.Response {
	/*
		oid: uint   当前object的ID
		pn: int   页数
	*/

	// 先进行类型转换
	oid, pn, err := TypeChange(oidStr, pnStr)
	if err != nil {
		return &serializer.Response{
			Code:  serializer.UserInputError,
			Msg:   "用户输入有误",
			Error: err.Error(),
		}
	}

	// 初始化接收数据库的变量
	var AllRootDB []model.Comment

	// 获取当前OID下的所有根评论
	if err := model.DB.Model(&model.Comment{}).Where("root = ? and object_id = ?", true, oid).Offset(PageSize * (pn - 1)).Limit(PageSize).Find(&AllRootDB).Error; err != nil {
		return &serializer.Response{
			Code:  serializer.DatabaseReadError,
			Msg:   "读取数据库失败",
			Error: err.Error(),
		}
	}

	// 创建data字段要显示的列表
	result := make([]serializer.CommentSerilizerResponse, 0)

	// 获取每一条根评论下面的所有子评论
	for _, RootDB := range AllRootDB {
		var ChildDB []model.Comment

		if err := model.DB.Model(&model.Comment{}).Where("root_id = ?", RootDB.ID).Find(&ChildDB).Error; err != nil {
			return &serializer.Response{
				Code:  serializer.DatabaseReadError,
				Msg:   "数据库读取失败",
				Error: err.Error(),
			}
		}

		count := 0
		if err := model.DB.Model(&model.Comment{}).Where("root_id = ?", RootDB.ID).Count(&count).Error; err != nil {
			return &serializer.Response{
				Code:  serializer.DatabaseReadError,
				Msg:   "数据库读取失败",
				Error: err.Error(),
			}
		}

		result = append(result, serializer.CommentsResponse(RootDB, ChildDB, count))
	}

	return &serializer.Response{
		Data:      result,
		Msg:       "数据获取成功",
	}
}
