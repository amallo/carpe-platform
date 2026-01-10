package infra

type FakeConfigGateway struct {
	deviceID string
}

func NewFakeConfigGateway() *FakeConfigGateway {
	return &FakeConfigGateway{}
}

func (g *FakeConfigGateway) GetDeviceID() string {
	return g.deviceID
}

func (g *FakeConfigGateway) SetDeviceID(deviceID string) {
	g.deviceID = deviceID
}
