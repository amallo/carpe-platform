package gateways

type ConfigGateway interface {
	GetDeviceID() (string, error)
	SetDeviceID(deviceID string) error
}
