package models

import "time"

type Room struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	Messages  []Message
}
