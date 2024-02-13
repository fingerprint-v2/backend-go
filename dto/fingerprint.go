package dto

type CreateFingerprintReq struct {
	Wifis []CreateWifiReq `json:"wifis" validate:"required,min=1,dive"`
}

type SearchFingerprintReq struct{}
