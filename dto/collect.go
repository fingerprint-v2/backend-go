package dto

type CreateSurveyReq struct {
	PointID       string                 `json:"point_id" validate:"required,uuid4"`
	CollectDevice CreateCollectDeviceReq `json:"collect_device"`
	Mode          string                 `json:"mode" validate:"required,oneof=SUPERVISED UNSUPERVISED PREDICTION"`
	SiteID        string                 `json:"site_id" validate:"omitempty,uuid4"`
}
