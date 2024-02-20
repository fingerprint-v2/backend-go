package dto

type CreateTrainingReq struct {
	SiteID       string `json:"site_id" validate:"required,uuid4"`
	TrainingName string `json:"training_name" validate:"required"`
	TrainingType string `json:"training_type" validate:"required,oneof=SUPERVISED UNUPERVISED"`
}
