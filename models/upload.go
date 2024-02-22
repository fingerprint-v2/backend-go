package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Upload struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	// Mode: SURVEY_SUPERVISED, SURVEY_UNSUPERVISED, PREDICTION_TRIAL, PREDICTION_TESTING, PREDICTION_TRACKING
	UploadMode string `json:"upload_mode" gorm:"type:varchar(255);not null"`

	ScanMode     string `json:"scan_mode" gorm:"type:varchar(255);not null"`
	ScanInterval *int   `json:"scan_interval" gorm:"type:integer"`
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
