package dto

type CreateFloorReq struct {
	Name       string  `json:"name" validate:"required"`
	Number     float32 `json:"number" validate:"required,numeric"`
	BuildingID string  `json:"building_id" validate:"required,uuid4"`
}

type SearchFloorReq struct {
	ID     string  `json:"id,omitempty" validate:"omitempty,uuid4"`
	Name   string  `json:"name,omitempty"`
	Number float32 `json:"number,omitempty"`
	//
	OrganizationID string `json:"organization_id,omitempty" validate:"omitempty,uuid4"`
	SiteID         string `json:"site_id,omitempty" validate:"omitempty,uuid4"`
	BuildingID     string `json:"building_id,omitempty" validate:"omitempty,uuid4"`
	//
	WithOrganization bool `json:"with_organization,omitempty"`
	WithSite         bool `json:"with_site,omitempty"`
	WithBuilding     bool `json:"with_building,omitempty"`
	//
	WithPoint bool `json:"with_points,omitempty"`
	All       bool `json:"all,omitempty"`
}

type DeleteFloorReq struct {
	ID string `json:"id" validate:"required,uuid4"`
}

type UpdateFloorReq struct {
	ID     string  `json:"id" validate:"required,uuid4"`
	Name   string  `json:"name"`
	Number float32 `json:"number" validate:"omitempty,numeric"`
	//
	BuildingID string `json:"building_id" validate:"omitempty,uuid4"`
}
