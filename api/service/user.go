package service

import (
	"fmt"
	"msg-board/api/domain"
	"msg-board/api/dto"
)

func (s *MessageService) SaveUser(req dto.UserRequest) error {

	m := domain.User{
		FirstName: req.FirstName,
		Email:     req.Email,
		Password:  req.Password,
	}

	err := s.repo.SaveUser(m)
	if err != nil {
		return err
	}
	return nil
}

func (s *MessageService) CheckPassword(req dto.UserRequest) (int, error) {

	m := domain.User{
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := s.repo.CheckPassword(m.Email)
	if err != nil {
		return 0, err
	}

	if user.Password != m.Password {
		return 0, fmt.Errorf("wrong password")
	}
	return user.ID, nil
}
