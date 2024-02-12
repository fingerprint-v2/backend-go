package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Fingerprint struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	// Mode: SUPERVISED, UNSUPERVISED, PREDICTION
	Mode              string `json:"mode" gorm:"type:varchar(255);not null"`
	IsOutsideCoverage bool   `json:"is_outside_coverage" gorm:"type:boolean;not null;default:false"`
	IsBetweenPoints   bool   `json:"is_between_points" gorm:"type:boolean;not null;default:false"`
	IsCurrent         bool   `json:"is_current" gorm:"type:boolean;not null;default:false"`
	//
	CollectDevice   CollectDevice `json:"collect_device" gorm:"foreignKey:CollectDeviceID;references:ID"`
	CollectDeviceID string        `json:"collect_device_id" gorm:"type:uuid;not null"`
	// Organization   Organization `json:"organization" gorm:"foreignKey:OrganizationID;references:ID"`
	// OrganizationID string       `json:"organization_id" gorm:"type:uuid;not null"`
	// Nullable
	Label       *Point       `json:"label" gorm:"foreignKey:LabelID;references:ID"`
	LabelID     *string      `json:"label_id" gorm:"type:uuid"`
	Predictions []Prediction `json:"predictions" gorm:"foreignKey:FingerprintID;references:ID"`
	NearPoints  []Point      `json:"near_points,omitempty" gorm:"many2many:fingerprint_near_points;"`
	Wifis       []Wifi       `json:"wifis,omitempty" gorm:"foreignKey:FingerprintID;references:ID"`
	Upload      Upload       `json:"upload,omitempty" gorm:"foreignKey:UploadID;references:ID"`
	UploadID    string       `json:"upload_id" gorm:"type:uuid;not null"`
	//
	TrackedEntityID   string `json:"tracked_entity" gorm:"type:varchar(255);not null"`
	TrackedEntityType string `json:"tracked_entity_type" gorm:"type:varchar(255);not null"`
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
