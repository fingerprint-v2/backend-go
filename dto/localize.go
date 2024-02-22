package dto

type CreateSurpervisedSurveyReq struct {
	PointLabelID  string                 `json:"point_label_id" validate:"required,uuid4"`
	CollectDevice CreateCollectDeviceReq `json:"collect_device"`
	ScanMode      string                 `json:"scan_mode" validate:"required,oneof=INTERVAL SINGLE"`
	ScanInterval  *int                   `json:"scan_interval" validate:"omitempty,numeric"`
	//
	IsBetweenPoints   bool `json:"is_between_points"`
	IsOutsideCoverage bool `json:"is_outside_coverage"`
	//
	Fingerprints []CreateFingerprintReq `json:"fingerprints" validate:"required,min=1,dive"`
}

type CreateUnsurpervisedSurveyReq struct {
	CollectDevice CreateCollectDeviceReq `json:"collect_device"`
	SiteID        string                 `json:"site_id" validate:"required,uuid4"`
	ScanMode      string                 `json:"scan_mode" validate:"required,oneof=INTERVAL SINGLE"`
	ScanInterval  *int                   `json:"scan_interval" validate:"omitempty,numeric"`
	//
	Fingerprints []CreateFingerprintReq `json:"fingerprints" validate:"required,min=1,dive"`
}

type CreatePredictionTestingReq struct {
	SiteID        string                 `json:"site_id" validate:"required,omitempty,uuid4"`
	CollectDevice CreateCollectDeviceReq `json:"collect_device"`
	Fingerprints  []CreateFingerprintReq `json:"fingerprints" validate:"required,min=1,dive"`
}

type CreatePredictionTrackingReq struct {
	UploadID                   string                 `json:"upload_id" validate:"required,uuid4"`
	CollectDevice              CreateCollectDeviceReq `json:"collect_device"`
	ExternalEntityExternalName string                 `json:"external_entity_external_name"`
	ExternalEntityExternalID   string                 `json:"external_entity_external_id" validate:"required"`
}
