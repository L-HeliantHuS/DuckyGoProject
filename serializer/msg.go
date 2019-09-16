package serializer

import (
	"DuckyGo/model"
)

// Msg 留言序列化器
type Msg struct {
	ID        uint   `json:"id"`
	Message   string `json:"message"`
	Subject   string `json:"subject"`
	UserID    uint   `json:"user_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// 响应序列化器
func MsgResponse(db model.Msg) Msg {
	return Msg{
		ID:        db.ID,
		Message:   db.Message,
		Subject:   db.Subject,
		UserID:    db.UserID,
		CreatedAt: db.CreatedAt.Unix(),
		UpdatedAt: db.UpdatedAt.Unix(),
	}
}

// OneMsg 单个留言序列化
type OneMsg struct {
	Message Msg `json:"msg"`
}

// OneMsg 多个留言序列化
type AnyMsg struct {
	Message []Msg `json:"msgs"`
	Count   int   `json:"count"`
}

// 响应
// 单个留言序列化响应
func OneMsgResponse(db model.Msg) OneMsg {
	return OneMsg{
		Message: MsgResponse(db),
	}
}

// 多个留言序列化响应
func MsgsResponse(db []model.Msg, count int) AnyMsg {
	var msg []Msg
	for _, i := range db {
		temp := MsgResponse(i)
		msg = append(msg, temp)
	}


	return AnyMsg{
		Message: msg,
		Count:   count,
	}

}
