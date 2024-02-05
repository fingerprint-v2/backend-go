package dto

type CreateSiteReq struct {
	Name           string `json:"name" validate:"required"`
	OrganizationID string `json:"organization_id" validate:"required,uuid4"`
}

// I have to use string as ID because zero-UUID is not considered empty. See https://github.com/upper/db/issues/624#issuecomment-1836279092
type SearchSiteReq struct {
	ID   string `json:"id,omitempty" validate:"omitempty,uuid4"`
	Name string `json:"name,omitempty" validate:"omitempty"`
	All  bool   `json:"all,omitempty" validate:"omitempty"`
}
