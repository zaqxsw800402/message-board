package domain

import (
	"msg-board/api/dto"
	"time"
)

type Message struct {
	Id        int `gorm:"primaryKey;autoIncrement"`
	Username  string
	Message   string
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (m Message) ToDto() dto.MessageResponse {
	return dto.MessageResponse{
		ID:       m.Id,
		Username: m.Username,
		Message:  m.Message,
		Time:     m.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
