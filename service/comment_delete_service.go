package service

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
	"errors"
	"strconv"
)

// DeleteCommentService 删除评论服务
type DeleteCommentService struct {
}

// CheckUserPermisson 检查当前用户有没有这条评论的操作权限
func CheckUserPermisson(id int, userID uint) error {
	var comment model.Comment
	model.DB.First(&comment, id)

	if comment.UserID != userID {
		return errors.New("The CurrentUser Not Permission Delete This Data.")
	}

	return nil
}

// Delete 删除评论
func (DeleteCommentService) Delete(idStr string, user *model.User) *serializer.Response {
	// 判断用户输入是否是真的数字
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return &serializer.Response{
			Code:  serializer.UserInputError,
			Msg:   "输入有误, 并不是合法的ID",
			Error: err.Error(),
		}
	}

	// 判断用户输入的这个oid是否是真实存在的
	num := 0
	if err := model.DB.Model(&model.Comment{}).Where("id = ?", id).Count(&num).Error; err != nil {
		return &serializer.Response{
			Code: serializer.UserInputError,
			Msg:  "这条信息不存在",
		}
	}

	// 检测这条留言是否是当前用户创建的
	if err := CheckUserPermisson(id, user.ID); err != nil {
		return &serializer.Response{
			Code:  serializer.UserNotPermissionError,
			Msg:   "你没有权限删除此条留言",
			Error: err.Error(),
		}
	}

	// 删除这条oid的内容
	if err := model.DB.Where("id = ?", id).Delete(&model.Comment{}).Error; err != nil {
		return &serializer.Response{
			Code:  serializer.DatabaseDeleteError,
			Msg:   "数据库删除失败",
			Error: err.Error(),
		}
	}

	// 删除和这条oid相关的内容
	if err := model.DB.Where("reply_id = ? or root_id = ?", id, id).Delete(&model.Comment{}).Error; err != nil {
		return &serializer.Response{
			Code:  serializer.DatabaseDeleteError,
			Msg:   "数据库删除失败.",
			Error: err.Error(),
		}
	}

	return &serializer.Response{
		Msg: "删除成功",
	}
}
