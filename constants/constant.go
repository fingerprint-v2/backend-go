package constants

import "time"

type UserRole int
type UploadMode int
type ScanMode int
type ExternalEntityType int

const (

	// JWTExpiration represents 24 hours in seconds
	JWTExpiration = time.Hour * 24
)

const (
	SUPERADMIN UserRole = iota + 1
	ADMIN
	USER
)

const (
	INTERVAL ScanMode = iota + 1
	SINGLE
)

const (
	SURVEY_SUPERVISED UploadMode = iota + 1
	SURVEY_UNSUPERVISED
	PREDICTION_TRIAL
	PREDICTION_TESTING
	PREDICTION_TRACKING
)

const (
	USER_MOBILE ExternalEntityType = iota + 1
	EMBEDDED_DEVICE
)

//go:generate stringer -type=UserRole,UploadMode,ScanMode,ExternalEntityType -output=constant_string.go
