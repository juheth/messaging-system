package message

import (
	"github.com/juheth/messaging-system/internal/domain"
)

type Service struct {
	messageRepo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		messageRepo: repo,
	}
}

func (s *Service) CreateMessage(content string, sender domain.User, roomID int) (*domain.Message, error) {
	message := domain.NewMessage(content, sender, roomID)
	err := s.messageRepo.Save(message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (s *Service) GetMessagesByRoom(roomID int) ([]domain.Message, error) {
	return s.messageRepo.GetByRoomID(roomID)
}

func (s *Service) UpdateMessage(id int, content string) (*domain.Message, error) {
	message, err := s.messageRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	message.Content = content
	return s.messageRepo.Update(message)
}

func (s *Service) DeleteMessage(id int) error {
	return s.messageRepo.Delete(id)
}
