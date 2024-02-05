package dto

type CreateUserReq struct {
	Username       string `json:"username" validate:"required"`
	Password       string `json:"password" validate:"required"`
	Role           string `json:"role" validate:"required"`
	OrganizationID string `json:"organization_id" validate:"required,uuid4"`
}

type UpdateUserReq struct {
	ID             string `json:"id" validate:"required,uuid4"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Role           string `json:"role"`
	OrganizationID string `json:"organization_id"`
}

type DeleteUserReq struct {
	ID string `json:"id" validate:"required,uuid4"`
}

type SearchUserReq struct {
	ID               string `json:"id,omitempty" validate:"omitempty,uuid4"`
	Username         string `json:"username,omitempty" validate:"omitempty"`
	Role             string `json:"role,omitempty" validate:"omitempty"`
	OrganizationID   string `json:"organization_id,omitempty" validate:"omitempty,uuid4"`
	WithOrganization bool   `json:"with_organization,omitempty" validate:"omitempty"`
	All              bool   `json:"all,omitempty" validate:"omitempty"`
}

type CookiePayloadUser struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	Role           string `json:"role"`
	OrganizationID string `json:"organization_id"`
}
