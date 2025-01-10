// infrastructure/database/connection.go
package database

import (
	"github.com/juheth/messaging-system/internal/domain"
	"gorm.io/gorm"
)

type GormMessageRepository struct {
	db *gorm.DB
}

func NewGormMessageRepository(db *gorm.DB) *GormMessageRepository {
	return &GormMessageRepository{db: db}
}

func (r *GormMessageRepository) Save(message *domain.Message) error {
	return r.db.Create(message).Error
}

func (r *GormMessageRepository) GetByRoomID(roomID int) ([]domain.Message, error) {
	var messages []domain.Message
	err := r.db.Where("room_id = ?", roomID).Find(&messages).Error
	return messages, err
}
