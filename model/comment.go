package model

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	ObjectID uint   // 回复的主题ID
	ReplyID  uint   // 回复的评论ID
	UserID   uint   // 用户ID
	Content  string // 评论内容
	RootID   uint   // 根评论ID
	Root     bool   // 是否是根评论
	Child    bool   // 是否是子评论
	LikeNum  uint   // 点赞数量
}
