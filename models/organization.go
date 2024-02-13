package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Organization struct {
	ID        uuid.UUID       `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name      string          `json:"name" gorm:"type:varchar(255);unique;not null"`
	IsSystem  bool            `json:"is_system" gorm:"default:false"`
	CreatedAt time.Time       `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time      `json:"updated_at" gorm:"<-:update"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	//
	Users            []User           `json:"users"`
	ExternalEntities []ExternalEntity `json:"external_entities"`
	Sites            []Site           `json:"sites"`
	Buildings        []Building       `json:"buildings"`
	Floors           []Floor          `json:"floors"`
	Points           []Point          `json:"points"`
	Fingerprints     []Fingerprint    `json:"fingerprints"`
}

// Internal search
// I have to use string as ID because zero-UUID is not considered empty and will mess up the search. See https://github.com/upper/db/issues/624#issuecomment-1836279092
type OrganizationFind struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	IsSystem bool   `json:"is_system,omitempty"`
}
