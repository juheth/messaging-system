package message

import (
	"github.com/juheth/messaging-system/internal/domain"
)

type Repository interface {
	Save(message *domain.Message) error
	GetByRoomID(roomID int) ([]domain.Message, error)
	GetByID(id int) (*domain.Message, error)
	Update(message *domain.Message) (*domain.Message, error)
	Delete(id int) error
}
