//go:build !tinygo

package infra

type FakeDeviceIdGenerator struct {
	deviceID string
	err      error
}

func NewFakeDeviceIdGenerator() *FakeDeviceIdGenerator {
	return &FakeDeviceIdGenerator{}
}

func (g *FakeDeviceIdGenerator) GenerateDeviceID() (string, error) {
	if g.err != nil {
		return "", g.err
	}
	return g.deviceID, nil
}

func (g *FakeDeviceIdGenerator) WillGenerateDeviceID(deviceID string) {
	g.deviceID = deviceID
	g.err = nil
}

func (g *FakeDeviceIdGenerator) WillFailWithError(err error) {
	g.err = err
	g.deviceID = ""
}
