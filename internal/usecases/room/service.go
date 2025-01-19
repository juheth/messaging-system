package room

import (
	"github.com/juheth/messaging-system/internal/domain"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateRoom(name string) (*domain.Room, error) {
	rm := &domain.Room{
		Name: name,
	}
	if err := s.repo.CreateRoom(rm); err != nil {
		return nil, err
	}
	return rm, nil
}

func (s *Service) UpdateRoom(id int, name string) (*domain.Room, error) {
	rm, err := s.repo.GetRoomByID(id)
	if err != nil {
		return nil, err
	}
	rm.Name = name
	if err := s.repo.UpdateRoom(rm); err != nil {
		return nil, err
	}
	return rm, nil
}

func (s *Service) DeleteRoom(id int) error {
	return s.repo.DeleteRoom(id)
}
