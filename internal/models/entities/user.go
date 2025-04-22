package entities

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
// @Description User represents a user in the system
type User struct {
	ID           int64          `gorm:"primaryKey" json:"id"`
	Username     string         `gorm:"unique;not null" json:"username"`
	Email        string         `gorm:"unique;not null" json:"email"`
	Password     string         `gorm:"not null" json:"password"`
	RefreshToken string         `json:"refresh_token"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-" swaggerignore:"true"`
}
