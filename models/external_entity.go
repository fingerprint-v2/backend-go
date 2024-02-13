package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Use pointer for time field https://stackoverflow.com/a/32646035
type ExternalEntity struct {
	ID        uuid.UUID       `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name      string          `json:"name" gorm:"type:varchar(255);unique;not null"`
	Type      string          `json:"type" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time       `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time      `json:"updated_at" gorm:"<-:update"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	//
	Organization   *Organization `json:"organization" gorm:"foreignKey:OrganizationID;references:ID"`
	OrganizationID string        `json:"organization_id" gorm:"type:uuid;not null"`
	//
	Fingerprints []Fingerprint `json:"fingerprints" gorm:"polymorphic:TrackedEntity;polymorphic_value:external_entity"`
}

// Internal search
type ExternalEntityFind struct {
	ID string `json:"id,omitempty"`
}
