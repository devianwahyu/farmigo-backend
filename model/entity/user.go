package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID       uint           `json:"user_id" gorm:"primaryKey;autoIncrement"`
	RoleID       uint           `json:"role_id"`
	Role         Role           ``
	Name         string         `json:"name" gorm:"not null;size:100"`
	Email        string         `json:"email" gorm:"not null;size:100;unique"`
	Password     string         `json:"password" gorm:"not null;size:100"`
	Gender       string         `json:"gender" gorm:"size:6"`
	PhoneNumber  string         `json:"phone_number" gorm:"size:15"`
	Address      string         `json:"address" gorm:"size:256"`
	Description  string         `json:"description" gorm:"size:256"`
	ProfileImage string         `json:"profile_image" gorm:"size:256"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
