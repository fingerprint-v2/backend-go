package dto

type CreateOrganizationReq struct {
	Name string `json:"name" validate:"required"`
}

type UpdateOrganizationReq struct {
	Name string `json:"name" validate:"required"`
}

// I have to use string as ID because zero-UUID is not considered empty. See https://github.com/upper/db/issues/624#issuecomment-1836279092
type SearchOrganizationReq struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
