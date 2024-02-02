package constants

import "time"

type UserRole int

const (

	// JWTExpiration represents 24 hours in seconds
	JWTExpiration = time.Hour * 24
)

const (
	SUPERADMIN UserRole = iota + 1
	ADMIN
	USER
)

//go:generate stringer -type=UserRole
