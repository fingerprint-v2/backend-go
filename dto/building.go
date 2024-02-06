package dto

type CreateBuildingReq struct {
	Name         string `json:"name" validate:"required"`
	ExternalName string `json:"external_name"`
	SiteID       string `json:"site_id" validate:"required,uuid4"`
}

type SearchBuildingReq struct {
	ID           string `json:"id,omitempty" validate:"omitempty,uuid4"`
	Name         string `json:"name,omitempty"`
	ExternalName string `json:"external_name,omitempty"`
	//
	OrganizationID string `json:"organization_id,omitempty" validate:"omitempty,uuid4"`
	SiteID         string `json:"site_id,omitempty" validate:"omitempty,uuid4"`
	//
	WithOrganization bool `json:"with_organization,omitempty"`
	WithSite         bool `json:"with_site,omitempty"`
	//
	WithFloor bool `json:"with_floor,omitempty"`
	WithPoint bool `json:"with_point,omitempty"`
	All       bool `json:"all,omitempty"`
}

type DeleteBuildingReq struct {
	ID string `json:"id" validate:"required,uuid4"`
}

type UpdateBuildingReq struct {
	ID           string `json:"id" validate:"required,uuid4"`
	Name         string `json:"name"`
	ExternalName string `json:"external_name"`
	//
	SiteID string `json:"site_id" validate:"omitempty,uuid4"`
}
