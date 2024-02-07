package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Wifi struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	SSID         string    `json:"ssid" gorm:"type:varchar(255);not null"`
	BSSID        string    `json:"bssid" gorm:"type:varchar(255);not null"`
	Capabilities string    `json:"capabilities" gorm:"type:varchar(255);not null"`
	Frequency    int       `json:"frequency" gorm:"type:integer;not null"`
	Level        int       `json:"level" gorm:"type:integer;not null"`
	//
	Fingerprint   Fingerprint `json:"fingerprint" gorm:"foreignKey:FingerprintID;references:ID"`
	FingerprintID string      `json:"fingerprint_id:" gorm:"type:uuid;not null"`
	//
	CreatedAt time.Time       `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time      `json:"updated_at" gorm:"<-:update"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// Internal search
type WifiFind struct {
	ID    string `json:"id,omitempty"`
	SSID  string `json:"ssid"`
	BSSID string `json:"bssid"`
}
