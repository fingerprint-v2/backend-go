package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CollectDevice struct {
	ID                 uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	DeviceUID          string    `json:"device_uid" gorm:"type:varchar(255);unique;not null"` // Unique
	DeviceID           string    `json:"device_id" gorm:"type:varchar(255);not null"`
	DeviceCarrier      string    `json:"device_carrier" gorm:"type:varchar(255);not null"`
	DeviceManufacturer string    `json:"device_manufacturer" gorm:"type:varchar(255);not null"`
	DeviceModel        string    `json:"device_model" gorm:"type:varchar(255);not null"`
	//
	Fingerprints []Fingerprint `json:"fingerprints" gorm:"foreignKey:CollectDeviceID;references:ID"`
	//
	CreatedAt time.Time       `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time      `json:"updated_at" gorm:"<-:update"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Internal search
type CollectDeviceFind struct {
	ID                 string `json:"id,omitempty"`
	DeviceUID          string `json:"device_uid,omitempty"`
	DeviceID           string `json:"device_id,omitempty"`
	DeviceCarrier      string `json:"device_carrier,omitempty"`
	DeviceManufacturer string `json:"device_manufacturer,omitempty"`
	DeviceModel        string `json:"device_model,omitempty"`
	//
}
