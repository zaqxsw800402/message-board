package domain

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBClient(dsn string) (*gorm.DB, error) {
	// 讀取環境變數
	// 建立與資料庫的連結
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	//建立表格，如果沒有表格
	err = db.AutoMigrate(&Message{}, &User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
