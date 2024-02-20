package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name string    `json:"name" gorm:"type:varchar(255);not null"`
	Path string    `json:"path" gorm:"type:varchar(255);not null"`
	//
	Organization   *Organization `json:"organization" gorm:"foreignKey:OrganizationID;references:ID"`
	OrganizationID string        `json:"organization_id" gorm:"type:uuid;not null"`
	Site           *Site         `json:"site" gorm:"foreignKey:SiteID;references:ID"`
	SiteID         string        `json:"site_id" gorm:"type:uuid;not null"`
	//
	Predictions []Prediction `json:"predictions" gorm:"foreignKey:ModelID;references:ID"`
	//
	CreatedAt time.Time       `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time      `json:"updated_at" gorm:"<-:update"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
