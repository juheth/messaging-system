package domain

type Room struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Users    []User    `gorm:"many2many:room_users;" json:"users"`
	Messages []Message `gorm:"foreignKey:RoomID" json:"messages"`
}

func NewRoom(name string) *Room {
	return &Room{
		Name: name,
	}
}
