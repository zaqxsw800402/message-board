package domain

import (
	"fmt"
	"gorm.io/gorm"
)

type MessageRepository struct {
	client *gorm.DB
}

func NewMessageRepository(client *gorm.DB) *MessageRepository {
	return &MessageRepository{client: client}
}

func (r MessageRepository) SaveMessage(m Message) error {
	result := r.client.Create(&m)
	if err := result.Error; err != nil {
		return fmt.Errorf("failed to save message")
	}
	return nil
}

func (r MessageRepository) DeleteMessage(id string) error {
	result := r.client.Where("id", id).Delete(&Message{})
	if err := result.Error; err != nil {
		return fmt.Errorf("failed to delete message %s", err.Error())
	}
	return nil
}

func (r MessageRepository) UpdateMessage(m Message) error {
	result := r.client.Model(&m).Where("id", m.Id).Updates(
		Message{Username: m.Username, Message: m.Message, CreatedAt: m.CreatedAt})

	if err := result.Error; err != nil {
		return fmt.Errorf("failed to update message %s", err.Error())
	}
	return nil
}

func (r MessageRepository) Delete() {

}
func (r MessageRepository) GetMessages() ([]Message, error) {
	var m []Message

	result := r.client.Order("id").Find(&m)

	if err := result.Error; err != nil {
		return nil, fmt.Errorf("no message left")

	}
	return m, nil
}
