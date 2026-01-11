//go:build !tinygo

package infra

type FakeConfigGateway struct {
	deviceID string
}

func NewFakeConfigGateway() *FakeConfigGateway {
	return &FakeConfigGateway{}
}

func (g *FakeConfigGateway) GetDeviceID() (string, error) {
	return g.deviceID, nil
}

func (g *FakeConfigGateway) SetDeviceID(deviceID string) error {
	g.deviceID = deviceID
	return nil
}
