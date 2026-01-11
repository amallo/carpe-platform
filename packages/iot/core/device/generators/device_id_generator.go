package generators

type DeviceIdGenerator interface {
	GenerateDeviceID() (string, error)
}
