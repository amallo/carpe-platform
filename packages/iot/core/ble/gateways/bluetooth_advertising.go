package gateways

// BluetoothAdvertising définit l'interface pour démarrer l'advertising Bluetooth
type BluetoothAdvertising interface {
	Start() error
}
