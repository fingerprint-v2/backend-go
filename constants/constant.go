package constants

import "time"

type UserRole int
type CollectMode int
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

const (
	SUPERVISED CollectMode = iota + 1
	UNSUPERVISED
	PREDICTION
)

const (
	INTERVAL ScanMode = iota + 1
	SINGLE
)

//go:generate stringer -type=UserRole,CollectMode,ScanMode
