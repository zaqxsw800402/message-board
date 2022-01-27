package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Session struct {
	Token string `gorm:"column:token;primaryKey"`
	Data  []byte `gorm:"data;not null"`
	//Expiry int64 `gorm:"autoCreateTime"`
	Expiry time.Time `gorm:"column:expiry;not null"`
}

func GetDBClient(dsn string) (*gorm.DB, error) {
	// 讀取環境變數
	// 建立與資料庫的連結
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	sqlDB, _ := db.DB()

	// 最多閒置數量
	sqlDB.SetMaxIdleConns(10)
	// 最多連接數量
	sqlDB.SetMaxOpenConns(10)
	// 等待醉酒時間
	sqlDB.SetConnMaxIdleTime(time.Hour)

	if err != nil {
		return nil, err
	}

	return db, nil
}
