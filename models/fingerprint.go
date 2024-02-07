package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Fingerprint struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	IsUnsupervised bool      `json:"is_unsupervised" gorm:"type:boolean;not null;default:false"`
	//
	Device   Device `json:"device" gorm:"foreignKey:DeviceID;references:ID"`
	DeviceID string `json:"device_id" gorm:"type:uuid;not null"`
	User     User   `json:"user" gorm:"foreignKey:UserID;references:ID"`
	UserID   string `json:"user_id" gorm:"type:uuid;not null"`
	// Nullable
	Point   *Point  `json:"point" gorm:"foreignKey:PointID;references:ID"`
	PointID *string `json:"point_id" gorm:"type:uuid"`
	//
	CreatedAt time.Time       `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time      `json:"updated_at" gorm:"<-:update"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// Internal search
type FingerprintFind struct {
	ID             string `json:"id,omitempty"`
	IsUnsupervised bool   `json:"is_unsupervised" gorm:"type:boolean;not null;default:false"`
}
