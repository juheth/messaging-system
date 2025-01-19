package room

import (
	"github.com/juheth/messaging-system/internal/domain"
)

type Repository interface {
	CreateRoom(room *domain.Room) error
	GetRoomByID(id int) (*domain.Room, error)
	UpdateRoom(room *domain.Room) error
	DeleteRoom(id int) error
}
