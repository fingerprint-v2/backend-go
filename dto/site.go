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
	WithBuildings bool `json:"with_buildings,omitempty"`
	WithFloors    bool `json:"with_floors,omitempty"`
	WithPoints    bool `json:"with_points,omitempty"`
	All           bool `json:"all,omitempty"`
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
