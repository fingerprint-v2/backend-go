package validates

import (
	"github.com/google/uuid"
)

type CreateOrganizationReq struct {
	Name string `json:"name" validate:"required"`
}

type UpdateOrganizationReq struct {
	Name string `json:"name" validate:"required"`
}

type SearchOrganizationReq struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
