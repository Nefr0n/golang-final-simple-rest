package entities

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	ID        int64          `gorm:"primaryKey" json:"id"`
	UserID    int64          `gorm:"not null;index" json:"user_id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Bio       string         `json:"bio"`
	AvatarURL string         `json:"avatar_url"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" swaggerignore:"true"`
}
