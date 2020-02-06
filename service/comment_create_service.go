package service

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
	"errors"
	"strings"
)

type CreateCommentService struct {
	ObjectID     uint   `form:"oid" json:"oid" binding:"required"`
	RootID       uint   `form:"root" json:"root"`
	ReplyID      uint   `form:"rid" json:"rid"`
	ReplyContent string `form:"content" json:"content"`
}

// SecCheck 参数安全性检查
func SecCheck(root uint, reply uint, content string) error {
	num := 0
	model.DB.Model(&model.Comment{}).Where("id = ?", root).Count(&num)
	if num < 0 {
		return errors.New("RootID不存在")
	}

	num = 0
	model.DB.Model(&model.Comment{}).Where("id = ?", reply).Count(&num)
	if num < 0 {
		return errors.New("ReplyID不存在")
	}
	replace := strings.Replace(content, " ", "", -1)
	if replace == "" {
		return errors.New("回复内容不能为空")
	}

	return nil
}

// Create 创建评论
func (service *CreateCommentService) Create(user *model.User) *serializer.Response {

	if err := SecCheck(service.RootID, service.ReplyID, service.ReplyContent); err != nil {
		return &serializer.Response{
			Code:  serializer.UserInputError,
			Msg:   "参数存在异常.",
			Error: err.Error(),
		}
	}

	db := model.Comment{
		ObjectID: service.ObjectID,
		ReplyID:  service.ReplyID,
		UserID:   user.ID,
		Content:  service.ReplyContent,
		RootID:   service.RootID,
		Root:     service.ReplyID == 0,
		Child:    service.ReplyID != 0,
		LikeNum:  0,
	}

	if err := model.DB.Create(&db).Error; err != nil {
		return &serializer.Response{
			Code:      serializer.DatabaseWriteError,
			Msg:       "数据库写入失败",
			Error:     err.Error(),
			TimeStamp: 0,
		}
	}

	return &serializer.Response{
		Data: serializer.CommentResponse(db),
		Msg:  "数据返回成功",
	}
}
