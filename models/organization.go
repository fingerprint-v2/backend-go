package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Organization struct {
	ID        uuid.UUID       `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name      string          `json:"name" gorm:"type:varchar(255);unique;not null"`
	CreatedAt time.Time       `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time      `json:"updated_at" gorm:"<-:update"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
	Users     []User          `json:"users,omitempty"`
	Sites     []Site          `json:"sites,omitempty"`
}

type SearchOrganization struct {
	ID   string `json:"id,omitempty" validate:"omitempty,uuid4"`
	Name string `json:"name,omitempty" validate:"omitempty"`
}
