package core

type Event[T any] struct {
	Type    EventType
	Payload T
}

type EventType string

const (
	PowerOn                      EventType = "power_on"
	DeviceReady                  EventType = "device_ready"
	DeviceIdGenerationFailed     EventType = "device_id_generation_failed"
	ConfigGatewayError           EventType = "config_gateway_error"
	BluetoothAdvertisingError   EventType = "bluetooth_advertising_error"
)
