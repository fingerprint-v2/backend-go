package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Floor struct {
	ID        uuid.UUID       `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name      string          `json:"name" gorm:"type:varchar(255);not null"`
	Number    float32         `json:"number" gorm:"type:float;not null"` // Numeric floor number used for ordering.  Can be decimal such floor 1.5.
	CreatedAt time.Time       `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time      `json:"updated_at" gorm:"<-:update"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
	//
	Site           Site         `json:"site" gorm:"foreignKey:SiteID;references:ID"`
	SiteID         string       `json:"site_id" gorm:"type:uuid;not null"`
	Building       Building     `json:"building" gorm:"foreignKey:BuildingID;references:ID"`
	BuildingID     string       `json:"building_id" gorm:"type:uuid;not null"`
	Organization   Organization `json:"organization" gorm:"foreignKey:OrganizationID;references:ID"`
	OrganizationID string       `json:"organization_id" gorm:"type:uuid;not null"`
	//
	Points []Point `json:"points,omitempty"`
}

// Internal search
type FloorFind struct {
	ID             string  `json:"id,omitempty"`
	Name           string  `json:"name,omitempty"`
	Number         float32 `json:"number,omitempty"`
	OrganizationID string  `json:"organization_id,omitempty"`
	SiteID         string  `json:"site_id,omitempty"`
	BuildingID     string  `json:"building_id,omitempty"`
}
