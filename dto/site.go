package dto

type CreateSiteReq struct {
	Name           string `json:"name" validate:"required"`
	OrganizationID string `json:"organization_id" validate:"required,uuid4"`
}

type SearchSiteReq struct {
	ID   string `json:"id,omitempty" validate:"omitempty,uuid4"`
	Name string `json:"name,omitempty"`
	//
	OrganizationID string `json:"organization_id,omitempty" validate:"omitempty,uuid4"`
	//
	WithOrganization bool `json:"with_organization,omitempty"`
	//
	WithBuilding bool `json:"with_building,omitempty"`
	WithFloor    bool `json:"with_floor,omitempty"`
	WithPoint    bool `json:"with_point,omitempty"`
	All          bool `json:"all,omitempty"`
}

type DeleteSiteReq struct {
	ID string `json:"id" validate:"required,uuid4"`
}

type UpdateSiteReq struct {
	ID   string `json:"id" validate:"required,uuid4"`
	Name string `json:"name"`
	//
	OrganizationID string `json:"organization_id" validate:"omitempty,uuid4"`
}
