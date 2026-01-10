package infra

type FakeDeviceIdGenerator struct {
	deviceID string
}

func NewFakeDeviceIdGenerator() *FakeDeviceIdGenerator {
	return &FakeDeviceIdGenerator{}
}

func (g *FakeDeviceIdGenerator) GenerateDeviceID() (string, error) {
	return g.deviceID, nil
}

func (g *FakeDeviceIdGenerator) WillGenerateDeviceID(deviceID string) {
	g.deviceID = deviceID
}
