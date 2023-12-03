package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Building struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name      string         `json:"name" gorm:"type:varchar(255);unique;not null"`
	Site      Site           `json:"site" gorm:"foreignKey:SiteID;references:ID"`
	SiteID    string         `json:"site_id" gorm:"type:uuid;not null"`
	Floors    []Floor        `json:"floors"`
	Points    []Point        `json:"points"`
	CreatedAt time.Time      `json:"created_at" gorm:"<-:create"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"<-:update"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
