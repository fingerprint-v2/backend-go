package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Use pointer for time field https://stackoverflow.com/a/32646035
type User struct {
	ID             uuid.UUID       `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Username       string          `json:"username" gorm:"type:varchar(255);unique;not null"`
	Password       string          `json:"password,omitempty" gorm:"type:varchar(255);not null"`
	Role           string          `json:"role" gorm:"type:varchar(255);not null;default:'USER'"`
	CreatedAt      time.Time       `json:"created_at" gorm:"<-:create"`
	UpdatedAt      *time.Time      `json:"updated_at" gorm:"<-:update"`
	DeletedAt      *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
	OrganizationID string          `json:"organization_id" gorm:"type:uuid;not null"`
	Organization   *Organization   `json:"organization,omitempty" gorm:"foreignKey:OrganizationID;references:ID"`
}

// Internal search
type UserFind struct {
	ID             string `json:"id,omitempty"`
	Username       string `json:"username,omitempty"`
	Role           string `json:"role,omitempty"`
	OrganizationID string `json:"organization_id,omitempty"`
}
