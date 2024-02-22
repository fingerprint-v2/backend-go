package dto

type CreateSurveyReq struct {
	PointLabelID  string                 `json:"point_label_id" validate:"omitempty,uuid4"`
	CollectDevice CreateCollectDeviceReq `json:"collect_device"`
	SiteID        string                 `json:"site_id" validate:"required_without=PointLabelID,omitempty,uuid4"`
	ScanMode      string                 `json:"scan_mode" validate:"required,oneof=INTERVAL SINGLE"`
	ScanInterval  *int                   `json:"scan_interval" validate:"omitempty,numeric"`
	//
	UploadMode        string `json:"upload_mode" validate:"required,oneof=SURVEY_SUPERVISED SURVEY_UNSUPERVISED"`
	IsBetweenPoints   bool   `json:"is_between_points"`
	IsOutsideCoverage bool   `json:"is_outside_coverage"`
	//
	Fingerprints []CreateFingerprintReq `json:"fingerprints" validate:"required,min=1,dive"`
}

type CreatePredictionReq struct {
	SiteID        string                 `json:"site_id" validate:"required_without=PointLabelID,omitempty,uuid4"`
	UploadMode    string                 `json:"upload_mode" validate:"required,oneof=PREDICTION_TRIAL PREDICTION_TESTING PREDICTION_TRACKING"`
	CollectDevice CreateCollectDeviceReq `json:"collect_device"`
	Fingerprints  []CreateFingerprintReq `json:"fingerprints" validate:"required,min=1,dive"`
}

type CreatePredictionValidationReq struct {
	UploadID          string `json:"upload_id" validate:"required,uuid4"`
	TrackedEntityName string `json:"tracked_entity_name" validate:"required"`
	UserID            string `json:"user_id" validate:"required_with=TrackedEntityName,uuid4"`
}
