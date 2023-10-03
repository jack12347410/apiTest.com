package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	// ID     uint   `gorm:"primary_key"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name,omitempty"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreateUserInput struct {
	Name  string `gorm:"type:varchar(255);not null"`
	Email string `gorm:"uniqueIndex;not null"`
}
