package dto

type CreateSurveyReq struct {
	PointID string                 `json:"point_id" validate:"required,uuid4"`
	Device  CreateCollectDeviceReq `json:"device" validate:"required"`
}
