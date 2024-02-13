package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Upload struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	//
	Fingerprints []Fingerprint `json:"fingerprints" gorm:"foreignKey:UploadID;references:ID"`
	//
	User   User   `json:"user" gorm:"foreignKey:UserID;references:ID"`
	UserID string `json:"user_id" gorm:"type:uuid;not null"`
	//
	CreatedAt time.Time       `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time      `json:"updated_at" gorm:"<-:update"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Internal search
type UploadFind struct {
	ID string `json:"id,omitempty"`
	//
}
