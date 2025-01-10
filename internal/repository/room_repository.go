package repository

import (
	"github.com/juheth/messaging-system/internal/domain"
	"gorm.io/gorm"
)

type RoomRepository interface {
	CreateRoom(room *domain.Room) error
	GetRoomByID(id int) (domain.Room, error)
	GetAllRooms() ([]domain.Room, error)
}

type roomRepo struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepo{db: db}
}

func (r *roomRepo) CreateRoom(room *domain.Room) error {
	return r.db.Create(room).Error
}

func (r *roomRepo) GetRoomByID(id int) (domain.Room, error) {
	var room domain.Room
	err := r.db.First(&room, id).Error
	return room, err
}

func (r *roomRepo) GetAllRooms() ([]domain.Room, error) {
	var rooms []domain.Room
	err := r.db.Find(&rooms).Error
	return rooms, err
}
