package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Fingerprint struct {
	ID                uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	IsOutsideCoverage bool      `json:"is_outside_coverage" gorm:"type:boolean;not null;default:false"`
	IsBetweenPoints   bool      `json:"is_between_points" gorm:"type:boolean;not null;default:false"`
	IsCurrent         bool      `json:"is_current" gorm:"type:boolean;not null;default:false"`
	// Collect device information
	CollectDevice   *CollectDevice `json:"collect_device" gorm:"foreignKey:CollectDeviceID;references:ID"`
	CollectDeviceID string         `json:"collect_device_id" gorm:"type:uuid;not null"`
	// Upload information
	Upload   *Upload `json:"upload" gorm:"foreignKey:UploadID;references:ID"`
	UploadID string  `json:"upload_id" gorm:"type:uuid;not null"`
	//
	Organization   *Organization `json:"organization" gorm:"foreignKey:OrganizationID;references:ID"`
	OrganizationID string        `json:"organization_id" gorm:"type:uuid;not null"`
	Site           *Site         `json:"site" gorm:"foreignKey:SiteID;references:ID"`
	SiteID         string        `json:"site_id" gorm:"type:uuid;not null"`
	PointLabel     *Point        `json:"point_label" gorm:"foreignKey:PointLabelID;references:ID"`
	PointLabelID   *string       `json:"point_label_id" gorm:"type:uuid"`
	//
	Predictions []Prediction `json:"predictions" gorm:"foreignKey:FingerprintID;references:ID"`
	NearPoints  []Point      `json:"near_points" gorm:"many2many:fingerprint_near_points;"`
	Wifis       []Wifi       `json:"wifis" gorm:"foreignKey:FingerprintID;references:ID"`
	//
	CreatedAt time.Time       `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time      `json:"updated_at" gorm:"<-:update"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Internal search
type FingerprintFind struct {
	ID             string `json:"id,omitempty"`
	IsUnsupervised bool   `json:"is_unsupervised"`
}
