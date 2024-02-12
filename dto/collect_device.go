package dto

type CreateCollectDeviceReq struct {
	DeviceUID          string `json:"device_uid" validate:"required"`
	DeviceID           string `json:"device_id"`
	DeviceCarrier      string `json:"device_carrier"`
	DeviceManufacturer string `json:"device_manufacturer"`
	DeviceModel        string `json:"device_model"`
}

type SearchCollectDeviceReq struct {
	DeviceUID          string `json:"device_uid" validate:"required"`
	DeviceID           string `json:"device_id"`
	DeviceCarrier      string `json:"device_carrier"`
	DeviceManufacturer string `json:"device_manufacturer"`
	DeviceModel        string `json:"device_model"`
}
