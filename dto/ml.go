package dto

type CreateTrainingReq struct {
	TrainingName string `json:"training_name" validate:"required"`
	TrainingType string `json:"training_type" validate:"required,oneof=SUPERVISED UNUPERVISED"`
}
