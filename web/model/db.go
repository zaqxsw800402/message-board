package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type DB struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *DB {
	return &DB{DB: db}
}

type Message struct {
	Id        int `gorm:"primaryKey;autoIncrement"`
	Username  string
	Message   string
	CreatedAt time.Time `gorm:"column:created_at"`
}

type User struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement"`
	FirstName string `gorm:"column:first_name"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
}

func (m *DB) Authenticate(email, password string) (*User, error) {

	var user User
	result := m.DB.Table("users").Where("email = ?", email).Find(&user)
	if err := result.Error; err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, fmt.Errorf("wrong password")
	}

	return &user, nil

}

func (m *DB) FindMessage(id string) (*Message, error) {

	var msg Message
	result := m.DB.Where("id = ?", id).Find(&msg)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &msg, nil

}
