package core

// DeviceStatus représente le statut du device
type DeviceStatus string

const (
	DeviceStatusInitializing DeviceStatus = "initializing"
	DeviceStatusReady        DeviceStatus = "ready"
	DeviceStatusError        DeviceStatus = "error"
)

// State représente l'état global de l'application
type State struct {
	Status DeviceStatus
}

// NewState crée un nouvel état initial
func NewState() *State {
	return &State{
		Status: DeviceStatusInitializing,
	}
}

