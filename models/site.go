package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Site struct {
	ID             uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name           string         `json:"name" gorm:"type:varchar(255);unique;not null"`
	Organization   Organization   `json:"organization" gorm:"foreignKey:OrganizationID;references:ID"`
	OrganizationID string         `json:"organization_id" gorm:"type:uuid;not null"`
	Buildings      []Building     `json:"buildings"`
	Floors         []Floor        `json:"floors"`
	Points         []Point        `json:"points"`
	CreatedAt      time.Time      `json:"created_at" gorm:"<-:create"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"<-:update"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}
