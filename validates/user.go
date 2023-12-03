package validates

import "github.com/fingerprint/models"

type CreateUserReq struct {
	Username       string `json:"username" validate:"required"`
	Password       string `json:"password" validate:"required"`
	Role           string `json:"role"`
	OrganizationID string `json:"organization_id" validate:"required"`
}

type UpdateUserReq struct {
	models.User
	ID string `json:"id" validate:"required"`
}
