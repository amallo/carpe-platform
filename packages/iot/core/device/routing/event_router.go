package routing

import (
	"github.com/carpe-platform/iot-golang/core"
	"github.com/carpe-platform/iot-golang/core/device/commands"
)

// RouteEvent route un événement vers la commande appropriée du module device
func RouteEvent(event core.Event[any], deps *core.Dependencies) core.Command {
	switch event.Type {
	case core.PowerOn:
		return commands.NewSetupDeviceCommand(deps)
	default:
		return nil
	}
}
