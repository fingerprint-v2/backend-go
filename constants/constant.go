package constants

import "time"

type UserRole int
type UploadMode int
type ScanMode int

const (

	// JWTExpiration represents 24 hours in seconds
	JWTExpiration = time.Hour * 24
)

const (
	SUPERADMIN UserRole = iota + 1
	ADMIN
	USER
)

// Mode: SURVEY_SUPERVISED, SURVEY_UNSUPERVISED, PREDICTION_TRIAL, PREDICTION_TESTING, PREDICTION_TRACKING
const (
	SURVEY_SUPERVISED UploadMode = iota + 1
	SURVEY_UNSUPERVISED
	PREDICTION_TRIAL
	PREDICTION_TESTING
	PREDICTION_TRACKING
)

const (
	INTERVAL ScanMode = iota + 1
	SINGLE
)

//go:generate stringer -type=UserRole,UploadMode,ScanMode
