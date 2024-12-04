package models

import "time"

type Contract struct {
	ID        uint `gorm:"primaryKey"`
	Value     int64
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
