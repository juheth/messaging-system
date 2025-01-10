package room

import (
	"errors"

	"github.com/juheth/messaging-system/internal/domain"
	"github.com/juheth/messaging-system/internal/repository"
)

type Service struct {
	roomRepo repository.RoomRepository
}

func NewService(roomRepo repository.RoomRepository) *Service {
	return &Service{roomRepo: roomRepo}
}

func (s *Service) CreateRoom(name string) (domain.Room, error) {
	if name == "" {
		return domain.Room{}, errors.New("room name cannot be empty")
	}
	room := domain.Room{Name: name}
	if err := s.roomRepo.CreateRoom(&room); err != nil {
		return domain.Room{}, err
	}
	return room, nil
}
