package core

// State représente l'état global de l'application
type State struct {
	DeviceID string
}

// NewState crée un nouvel état initial
func NewState() *State {
	return &State{
		DeviceID: "",
	}
}

