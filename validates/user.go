package validates

type CreateUserReq struct {
	Username       string `json:"username" validate:"required"`
	Password       string `json:"password" validate:"required"`
	Role           string `json:"role"`
	OrganizationID string `json:"organization_id" validate:"required"`
}

type UpdateUserReq struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	Role           string `json:"role"`
	OrganizationID string `json:"organization_id"`
}

type SearchUserReq struct {
	ID             string `json:"id,omitempty"`
	Username       string `json:"username,omitempty"`
	Role           string `json:"role,omitempty"`
	OrganizationID string `json:"organization_id,omitempty"`
}

type CookiePayloadUser struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	Role           string `json:"role"`
	OrganizationID string `json:"organization_id"`
}
