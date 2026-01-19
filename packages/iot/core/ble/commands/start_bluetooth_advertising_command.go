package commands

import (
	"github.com/carpe-platform/iot-golang/core"
	ble_gateways "github.com/carpe-platform/iot-golang/core/ble/gateways"
)

type StartBluetoothAdvertisingCommand struct {
	bluetoothAdvertising ble_gateways.BluetoothAdvertising
}

func NewStartBluetoothAdvertisingCommand(bluetoothAdvertising ble_gateways.BluetoothAdvertising) *StartBluetoothAdvertisingCommand {
	return &StartBluetoothAdvertisingCommand{
		bluetoothAdvertising: bluetoothAdvertising,
	}
}

func (c *StartBluetoothAdvertisingCommand) Execute() []core.Event[any] {
	err := c.bluetoothAdvertising.Start()
	if err != nil {
		// En cas d'erreur, on pourrait émettre un événement d'erreur
		// Pour l'instant, on retourne juste un événement vide ou un événement d'erreur
		return []core.Event[any]{
			{Type: core.BluetoothAdvertisingError, Payload: err.Error()},
		}
	}
	// Pas besoin d'émettre d'événement de succès pour l'instant
	// L'advertising a démarré, c'est un effet de bord
	return []core.Event[any]{}
}
