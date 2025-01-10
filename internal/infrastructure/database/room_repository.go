package database

import (
	"github.com/juheth/messaging-system/internal/domain"
	"gorm.io/gorm"
)

type RoomRepository struct {
	db *gorm.DB
}

func (r *RoomRepository) CreateRoom(room *domain.Room) error {
	return r.db.Create(room).Error
}

func (r *RoomRepository) GetAllRooms() ([]domain.Room, error) {
	var rooms []domain.Room
	err := r.db.Find(&rooms).Error
	return rooms, err
}

func (r *RoomRepository) GetRoomByID(id int) (domain.Room, error) {
	var room domain.Room
	err := r.db.First(&room, id).Error
	return room, err
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{db: db}
}
