package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Point struct {
	ID         uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name       string         `json:"name" gorm:"type:varchar(255);unique;not null"`
	Site       Site           `json:"site" gorm:"foreignKey:SiteID;references:ID"`
	SiteID     string         `json:"site_id" gorm:"type:uuid;not null"`
	Building   Building       `json:"building" gorm:"foreignKey:BuildingID;references:ID"`
	BuildingID string         `json:"building_id" gorm:"type:uuid;not null"`
	Floor      Floor          `json:"floor" gorm:"foreignKey:FloorID;references:ID"`
	FloorID    string         `json:"floor_id" gorm:"type:uuid;not null"`
	CreatedAt  time.Time      `json:"created_at" gorm:"<-:create"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"<-:update"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
