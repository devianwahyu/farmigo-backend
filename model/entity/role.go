package entity

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	RoleID    uint           `json:"role_id" gorm:"primaryKey;autoIncrement"`
	Type      string         `json:"type" gorm:"not null;size:10"`
	User      []User         ``
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
