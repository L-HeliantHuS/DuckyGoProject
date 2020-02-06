package serializer

import "DuckyGo/model"

type CommentSerilizer struct {
	ID       uint   `json:"id"`
	ObjectID uint   `json:"oid"` // 图片ID
	ReplyID  uint   `json:"rid"` // 回复的ID
	UserID   uint   `json:"uid"` // 用户ID
	RootID   uint   `json:"root"`
	IsRoot   bool   `json:"is_root"`
	IsChild  bool   `json:"is_child"`
	Content  string `json:"content"`
	Time     int64  `json:"time"`
	LikeNum  uint   `json:"like"`
	Person   User   `json:"person"`
}

type CommentSerilizerResponse struct {
	ID       uint               `json:"id"`
	ObjectID uint               `json:"oid"` // 图片ID
	ReplyID  uint               `json:"rid"` // 回复的ID
	UserID   uint               `json:"uid"` // 用户ID
	RootID   uint               `json:"root"`
	IsRoot   bool               `json:"is_root"`
	IsChild  bool               `json:"is_child"`
	Content  string             `json:"content"`
	Time     int64              `json:"time"`
	LikeNum  uint               `json:"like"`
	Person   User               `json:"person"`
	Count    int                `json:"count"`
	Members  []CommentSerilizer `json:"members"`
}

// CommentResponse 评论结束后返回的响应
func CommentResponse(db model.Comment) CommentSerilizer {
	user, _ := model.GetUser(db.UserID)
	return CommentSerilizer{
		ID:       db.ID,
		ObjectID: db.ObjectID,
		ReplyID:  db.ReplyID,
		UserID:   db.UserID,
		RootID:   db.RootID,
		IsRoot:   db.Root,
		IsChild:  db.Child,
		Content:  db.Content,
		Time:     db.CreatedAt.Unix(),
		LikeNum:  db.LikeNum,
		Person:   BuildUser(user),
	}
}

// CommentsResponse 获取评论序列化器
func CommentsResponse(rootDB model.Comment, ChildDB []model.Comment, count int) CommentSerilizerResponse {
	var result []CommentSerilizer
	for _, i := range ChildDB {
		result = append(result, CommentResponse(i))
	}
	user, _ := model.GetUser(rootDB.UserID)
	return CommentSerilizerResponse{
		ID:       rootDB.ID,
		ObjectID: rootDB.ObjectID,
		ReplyID:  rootDB.ReplyID,
		UserID:   rootDB.UserID,
		RootID:   rootDB.RootID,
		IsRoot:   rootDB.Root,
		IsChild:  rootDB.Child,
		Content:  rootDB.Content,
		Time:     rootDB.CreatedAt.Unix(),
		LikeNum:  rootDB.LikeNum,
		Person:   BuildUser(user),
		Count:    count,
		Members:  result,
	}
}
