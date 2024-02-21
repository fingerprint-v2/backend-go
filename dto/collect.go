package dto

type CreateSurveyReq struct {
	PointLabelID  string                 `json:"point_label_id" validate:"omitempty,uuid4"`
	CollectDevice CreateCollectDeviceReq `json:"collect_device"`
	SiteID        string                 `json:"site_id" validate:"required_without=PointLabelID,omitempty,uuid4"`
	ScanMode      string                 `json:"scan_mode" validate:"required,oneof=INTERVAL SINGLE"`
	ScanInterval  *int                   `json:"scan_interval" validate:"omitempty,numeric"`
	//
	Mode              string `json:"mode" validate:"required,oneof=SUPERVISED UNSUPERVISED PREDICTION"`
	IsBetweenPoints   bool   `json:"is_between_points"`
	IsOutsideCoverage bool   `json:"is_outside_coverage"`
	//
	Fingerprints []CreateFingerprintReq `json:"fingerprints" validate:"required,min=1,dive"`
}
