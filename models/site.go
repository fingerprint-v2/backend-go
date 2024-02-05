package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Site struct {
	ID        uuid.UUID       `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name      string          `json:"name" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time       `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time      `json:"updated_at" gorm:"<-:update"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
	//
	Organization   *Organization `json:"organization" gorm:"foreignKey:OrganizationID;references:ID"`
	OrganizationID string        `json:"organization_id" gorm:"type:uuid;not null"`
	//
	Buildings []Building `json:"buildings,omitempty"`
	Floors    []Floor    `json:"floors,omitempty"`
	Points    []Point    `json:"points,omitempty"`
}

// Internal search
type SiteFind struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	OrganizationID string `json:"organization_id,omitempty"`
}
