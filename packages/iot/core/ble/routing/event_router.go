package routing

import (
	"github.com/carpe-platform/iot-golang/core"
	"github.com/carpe-platform/iot-golang/core/ble/commands"
)

// RouteEvent route un événement vers la commande appropriée du module BLE
func RouteEvent(event core.Event[any], deps *core.Dependencies) core.Command {
	switch event.Type {
	case core.DeviceReady:
		return commands.NewStartBluetoothAdvertisingCommand(deps.BluetoothAdvertising)
	default:
		return nil
	}
}
