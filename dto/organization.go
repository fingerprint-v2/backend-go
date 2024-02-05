package dto

type CreateOrganizationReq struct {
	Name     string `json:"name" validate:"required"`
	IsSystem bool   `json:"is_system" validate:"omitempty"`
}

type UpdateOrganizationReq struct {
	ID   string `json:"id" validate:"required,uuid4"`
	Name string `json:"name" validate:"required"`
}

type DeleteOrganizationReq struct {
	ID string `json:"id" validate:"required,uuid4"`
}

// I have to use string as ID because zero-UUID is not considered empty. See https://github.com/upper/db/issues/624#issuecomment-1836279092
type SearchOrganizationReq struct {
	ID            string `json:"id,omitempty" validate:"omitempty,uuid4"`
	Name          string `json:"name,omitempty" validate:"omitempty"`
	WithUsers     bool   `json:"with_users,omitempty" validate:"omitempty"`
	WithSites     bool   `json:"with_sites,omitempty" validate:"omitempty"`
	WithBuildings bool   `json:"with_buildings,omitempty" validate:"omitempty"`
	WithFloors    bool   `json:"with_floors,omitempty" validate:"omitempty"`
	WithPoints    bool   `json:"with_points,omitempty" validate:"omitempty"`
	All           bool   `json:"all,omitempty" validate:"omitempty"`
}
