package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Point struct {
	ID           uuid.UUID       `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name         string          `json:"name" gorm:"type:varchar(255);not null"`
	ExternalName string          `json:"external_name" gorm:"type:varchar(255)"`
	IsGroup      bool            `json:"is_group" gorm:"type:boolean;not null;default:false"`
	CreatedAt    time.Time       `json:"created_at" gorm:"<-:create"`
	UpdatedAt    *time.Time      `json:"updated_at" gorm:"<-:update"`
	DeletedAt    *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	//
	Floor          *Floor        `json:"floor" gorm:"foreignKey:FloorID;references:ID"`
	FloorID        string        `json:"floor_id" gorm:"type:uuid;not null"`
	Building       *Building     `json:"building" gorm:"foreignKey:BuildingID;references:ID"`
	BuildingID     string        `json:"building_id" gorm:"type:uuid;not null"`
	Site           *Site         `json:"site" gorm:"foreignKey:SiteID;references:ID"`
	SiteID         string        `json:"site_id" gorm:"type:uuid;not null"`
	Organization   *Organization `json:"organization" gorm:"foreignKey:OrganizationID;references:ID"`
	OrganizationID string        `json:"organization_id" gorm:"type:uuid;not null"`
	// Fingerprint Reference
	Fingerprints *[]Fingerprint `json:"fingerprints" gorm:"foreignKey:PointLabelID;references:ID"`
	// Prediction Reference
	Predictions *[]Prediction `json:"predictions" gorm:"foreignKey:PointID;references:ID"`
	// Self-referential: Grouping
	GroupID *string  `json:"group_id" gorm:"type:uuid"`
	Members *[]Point `json:"members" gorm:"foreignKey:GroupID;references:ID"`
	// Many-to-many: Vicinity Points
	VicinityPoints []*Point `json:"vicinity_points" gorm:"many2many:point_vicinity_points;"`
}

// Internal search
type PointFind struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	ExternalName string `json:"external_name,omitempty"`
	IsGroup      bool   `json:"is_group,omitempty"`
	//
	FloorID        string `json:"floor_id,omitempty"`
	SiteID         string `json:"site_id,omitempty"`
	BuildingID     string `json:"building_id,omitempty"`
	OrganizationID string `json:"organization_id,omitempty"`
}
