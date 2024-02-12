package dto

type SurveyReq struct {
	PointID string `json:"point_id" validate:"required,uuid4"`
	// Device  CollectDeviceReq `json:"device" validate:"required"`
}
