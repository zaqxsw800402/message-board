package service

import (
	"msg-board/api/domain"
	"msg-board/api/dto"
	"strconv"
	"time"
)

type MessageService struct {
	repo *domain.MessageRepository
}

func NewMessageService(repo *domain.MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) SaveMessage(req dto.MessageRequest) error {

	m := domain.Message{
		Username:  req.Username,
		Message:   req.Message,
		CreatedAt: time.Now(),
	}
	err := s.repo.SaveMessage(m)
	if err != nil {
		return err
	}
	return nil
}

func (s *MessageService) UpdateMessage(id string, req dto.MessageRequest) error {
	msgID, _ := strconv.Atoi(id)
	m := domain.Message{
		Id:        msgID,
		Username:  req.Username,
		Message:   req.Message,
		CreatedAt: time.Now(),
	}
	err := s.repo.UpdateMessage(m)
	if err != nil {
		return err
	}
	return nil
}

func (s *MessageService) DeleteMessage(id string) error {
	err := s.repo.DeleteMessage(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *MessageService) GetMessages() ([]dto.MessageResponse, error) {
	msgs, err := s.repo.GetMessages()
	if err != nil {
		return nil, err
	}

	resp := make([]dto.MessageResponse, 0)
	for _, msg := range msgs {
		resp = append(resp, msg.ToDto())
	}
	return resp, nil
}
