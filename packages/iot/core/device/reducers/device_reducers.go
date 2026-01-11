package reducers

import "github.com/carpe-platform/iot-golang/core"

// ReduceDeviceEvents réduit les événements liés au device
// Ce reducer filtre les événements qui concernent le device et modifie le state en conséquence
// Retourne le state modifié ou inchangé si l'événement ne concerne pas le device
func ReduceDeviceEvents(event core.Event[any], state *core.State) *core.State {
	switch event.Type {
	case core.DeviceReady:
		state.Status = core.DeviceStatusReady
		return state
	case core.DeviceIdGenerationFailed:
		state.Status = core.DeviceStatusErrorCannotGenerateDeviceId
		return state
	case core.ConfigGatewayError:
		state.Status = core.DeviceStatusErrorConfigGateway
		return state
	default:
		// Ne réagit pas à cet événement
		return state
	}
}

