package domain

import "time"

type Message struct {
	ID         int       `json:"id"`
	Content    string    `json:"content"`
	UserID     int       `json:"user_id"`
	Sender     User      `gorm:"foreignKey:UserID"`
	RoomID     int       `json:"room_id"`
	Created_at time.Time `json:"created"`
}

func NewMessage(content string, sender User, roomID int) *Message {
	return &Message{
		Content:    content,
		UserID:     sender.ID,
		RoomID:     roomID,
		Created_at: time.Now(),
	}
}
