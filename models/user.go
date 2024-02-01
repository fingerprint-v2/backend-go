package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID             uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Username       string         `json:"username" gorm:"type:varchar(255);unique;not null"`
	Password       string         `json:"-" gorm:"type:varchar(255);not null"`
	Role           string         `json:"role" gorm:"type:varchar(255);not null;default:'USER'"`
	Organization   *Organization  `json:"organization,omitempty" gorm:"foreignKey:OrganizationID;references:ID"`
	OrganizationID string         `json:"organization_id" gorm:"type:uuid;not null"`
	CreatedAt      time.Time      `json:"created_at" gorm:"<-:create"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"<-:update"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
