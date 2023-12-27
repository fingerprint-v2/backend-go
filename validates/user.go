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
