package database

import (
	"github.com/juheth/messaging-system/internal/domain"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) Save(message *domain.Message) error {
	return r.db.Create(message).Error
}

func (r *MessageRepository) GetByRoomID(roomID int) ([]domain.Message, error) {
	var messages []domain.Message
	err := r.db.Where("room_id = ?", roomID).Find(&messages).Error
	return messages, err
}

func (r *MessageRepository) GetByID(id int) (*domain.Message, error) {
	var message domain.Message
	err := r.db.First(&message, id).Error
	return &message, err
}

func (r *MessageRepository) Update(message *domain.Message) (*domain.Message, error) {
	err := r.db.Save(message).Error
	return message, err
}

func (r *MessageRepository) Delete(id int) error {
	return r.db.Delete(&domain.Message{}, id).Error
}
