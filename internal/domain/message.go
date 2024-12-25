package models

import "time"

type Message struct {
	ID        uint
	RoomID    uint
	UserID    uint
	Content   string
	CreatedAt time.Time
}
