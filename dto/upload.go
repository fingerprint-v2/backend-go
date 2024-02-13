package dto

type SearchUploadReq struct {
	ID string `json:"id" validate:"omitempty,uuid"`
}
