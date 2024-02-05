package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Building struct {
	ID           uuid.UUID       `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name         string          `json:"name" gorm:"type:varchar(255);not null"`
	ExternalName string          `json:"external_name" gorm:"type:varchar(255)"`
	CreatedAt    time.Time       `json:"created_at" gorm:"<-:create"`
	UpdatedAt    *time.Time      `json:"updated_at" gorm:"<-:update"`
	DeletedAt    *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
	//
	Site           *Site         `json:"site" gorm:"foreignKey:SiteID;references:ID"`
	SiteID         string        `json:"site_id" gorm:"type:uuid;not null"`
	Organization   *Organization `json:"organization" gorm:"foreignKey:OrganizationID;references:ID"`
	OrganizationID string        `json:"organization_id" gorm:"type:uuid;not null"`
	//
	Floors []Floor `json:"floors,omitempty"`
	Points []Point `json:"points,omitempty"`
}

type BuildingFind struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	OrganizationID string `json:"organization_id,omitempty"`
	SiteID         string `json:"site_id,omitempty"`
}
