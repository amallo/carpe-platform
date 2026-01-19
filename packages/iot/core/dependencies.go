package core

import (
	ble_gateways "github.com/carpe-platform/iot-golang/core/ble/gateways"
	device_gateways "github.com/carpe-platform/iot-golang/core/device/gateways"
	device_generators "github.com/carpe-platform/iot-golang/core/device/generators"
)

type Dependencies struct {
	ConfigGateway       device_gateways.ConfigGateway
	DeviceIdGenerator   device_generators.DeviceIdGenerator
	BluetoothAdvertising ble_gateways.BluetoothAdvertising
}
