package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Prediction struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Probability float64   `json:"probability" gorm:"type:float;not null"`
	//
	Point         Point       `json:"point" gorm:"foreignKey:PointID;references:ID"`
	PointID       string      `json:"point_id" gorm:"type:uuid;not null"`
	Fingerprint   Fingerprint `json:"fingerprint" gorm:"foreignKey:FingerprintID;references:ID"`
	FingerprintID string      `json:"fingerprint_id" gorm:"type:uuid;not null"`
	//
	CreatedAt time.Time       `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time      `json:"updated_at" gorm:"<-:update"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Internal search
type PredictionFind struct {
	ID string `json:"id,omitempty"`
	//
}
