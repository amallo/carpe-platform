package commands

import "github.com/carpe-platform/iot-golang/core"

type SetupDeviceCommand struct {
	dependencies *core.Dependencies
}

func NewSetupDeviceCommand(dependencies *core.Dependencies) *SetupDeviceCommand {
	return &SetupDeviceCommand{
		dependencies: dependencies,
	}
}

func (c *SetupDeviceCommand) Execute() []core.Event[any] {
	// Check if a deviceID already exists
	existingDeviceID, err := c.dependencies.ConfigGateway.GetDeviceID()
	if err != nil {
		return []core.Event[any]{
			{Type: core.ConfigGatewayError, Payload: err.Error()},
		}
	}

	if existingDeviceID != "" {
		// Device already configured: reuse existing deviceID
		return []core.Event[any]{
			{Type: core.DeviceReady, Payload: nil},
		}
	}

	// Device not configured: generate a new deviceID
	deviceID, err := c.dependencies.DeviceIdGenerator.GenerateDeviceID()
	if err != nil {
		return []core.Event[any]{
			{Type: core.DeviceIdGenerationFailed, Payload: err.Error()},
		}
	}

	err = c.dependencies.ConfigGateway.SetDeviceID(deviceID)
	if err != nil {
		return []core.Event[any]{
			{Type: core.ConfigGatewayError, Payload: err.Error()},
		}
	}

	return []core.Event[any]{
		{Type: core.DeviceReady, Payload: nil},
	}
}
