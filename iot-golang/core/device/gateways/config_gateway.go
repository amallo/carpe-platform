package gateways

type ConfigGateway interface {
	GetDeviceID() string
	SetDeviceID(deviceID string)
}

type ConfigGatewayImpl struct {
	deviceID string
}
