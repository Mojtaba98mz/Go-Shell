package models

import "time"

type Command struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"index"`
	Name      string    `gorm:"not null"`
	Count     int       `gorm:"not null"`
	Timestamp time.Time `gorm:"not null"`
}
