package dto

type CreateSurveyReq struct {
	PointID string                 `json:"point_id" validate:"required,uuid4"`
	Device  CreateCollectDeviceReq `json:"device"`
	Mode    string                 `json:"mode" validate:"required,oneof=SUPERVISED UNSUPERVISED PREDICTION"`
}
