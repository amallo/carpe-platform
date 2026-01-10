package core

import (
	device_gateways "github.com/carpe-platform/iot-golang/core/device/gateways"
	device_generators "github.com/carpe-platform/iot-golang/core/device/generators"
)

type Dependencies struct {
	ConfigGateway     device_gateways.ConfigGateway
	DeviceIdGenerator device_generators.DeviceIdGenerator
}
