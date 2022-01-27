package domain

import (
	"database/sql"
	"fmt"
)

func (r MessageRepository) SaveUser(u User) error {
	result := r.client.Create(&u)
	if err := result.Error; err != nil {
		return fmt.Errorf("failed to save user")
	}
	return nil
}

func (r MessageRepository) CheckPassword(email string) (*User, error) {
	var user User
	// 在account表格裡預載入交易紀錄的資料，並且讀取特定id的資料
	result := r.client.Table("users").Where("email = ?", email).Find(&user)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no matching email")
	}

	if err := result.Error; err != nil {
		//logger_zap.Error("Error while querying accounts table" + err.Error())
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found by email")
		}
		return nil, fmt.Errorf("unexpected database error when get user by email")

	}
	return &user, nil
}
