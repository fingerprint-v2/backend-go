package dto

type CreateWifiReq struct {
	SSID         string `json:"ssid" validate:"required"`
	BSSID        string `json:"bssid" validate:"required,mac"`
	Capabilities string `json:"capabilities"`
	Frequency    int    `json:"frequency" validate:"required"`
	Level        int    `json:"level" validate:"required"`
}

type SearchWifiReq struct{}
