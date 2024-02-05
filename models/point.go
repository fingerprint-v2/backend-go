package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Point struct {
	ID             uuid.UUID       `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name           string          `json:"name" gorm:"type:varchar(255);not null"`
	IsUnsupervised bool            `json:"is_supervised" gorm:"type:boolean;not null;default:false"`
	CreatedAt      time.Time       `json:"created_at" gorm:"<-:create"`
	UpdatedAt      *time.Time      `json:"updated_at" gorm:"<-:update"`
	DeletedAt      *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
	//
	Site           *Site         `json:"site" gorm:"foreignKey:SiteID;references:ID"`
	SiteID         string        `json:"site_id" gorm:"type:uuid;not null"`
	Building       *Building     `json:"building" gorm:"foreignKey:BuildingID;references:ID"`
	BuildingID     string        `json:"building_id" gorm:"type:uuid;not null"`
	Floor          *Floor        `json:"floor" gorm:"foreignKey:FloorID;references:ID"`
	FloorID        string        `json:"floor_id" gorm:"type:uuid;not null"`
	Organization   *Organization `json:"organization" gorm:"foreignKey:OrganizationID;references:ID"`
	OrganizationID string        `json:"organization_id" gorm:"type:uuid;not null"`
	// Self-referential
	GroupID *string `json:"parent_id" gorm:"type:uuid"`
	Members []Point `json:"members,omitempty" gorm:"foreignKey:GroupID;references:ID"`
}

// Internal search
type PointFind struct {
	ID             string  `json:"id,omitempty"`
	Name           string  `json:"name,omitempty"`
	Number         float32 `json:"number,omitempty"`
	OrganizationID string  `json:"organization_id,omitempty"`
	SiteID         string  `json:"site_id,omitempty"`
	BuildingID     string  `json:"building_id,omitempty"`
	FloorID        string  `json:"floor_id,omitempty"`
}
