package dto

type CreatePointReq struct {
	Name         string `json:"name" validate:"required"`
	ExternalName string `json:"external_name"`
	FloorID      string `json:"floor_id" validate:"required,uuid4"`
}

type SearchPointReq struct {
	ID           string `json:"id,omitempty" validate:"omitempty,uuid4"`
	Name         string `json:"name,omitempty"`
	ExternalName string `json:"external_name"`
	//
	OrganizationID string `json:"organization_id,omitempty" validate:"omitempty,uuid4"`
	SiteID         string `json:"site_id,omitempty" validate:"omitempty,uuid4"`
	BuildingID     string `json:"building_id,omitempty" validate:"omitempty,uuid4"`
	FloorID        string `json:"floor_id,omitempty" validate:"omitempty,uuid4"`
	//
	All bool `json:"all,omitempty"`
}

type DeletePointReq struct {
	ID string `json:"id" validate:"required,uuid4"`
}

type UpdatePointReq struct {
	ID           string `json:"id" validate:"required,uuid4"`
	Name         string `json:"name"`
	ExternalName string `json:"external_name"`
	//
	FloorID string `json:"floor_id" validate:"omitempty,uuid4"`
}
