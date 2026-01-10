package commands

import "github.com/carpe-platform/iot-golang/core"

type SetupDeviceCommand struct {
	dependencies *core.Dependencies
	state        *core.State
}

func NewSetupDeviceCommand(dependencies *core.Dependencies, state *core.State) *SetupDeviceCommand {
	return &SetupDeviceCommand{
		dependencies: dependencies,
		state:        state,
	}
}

func (c *SetupDeviceCommand) Execute() error {
	// Check if a deviceID already exists
	existingDeviceID, err := c.dependencies.ConfigGateway.GetDeviceID()
	if err != nil {
		return err
	}

	if existingDeviceID != "" {
		// Device already configured: reuse existing deviceID
		c.state.Status = core.DeviceStatusReady
		return nil
	}

	// Device not configured: generate a new deviceID
	deviceID, err := c.dependencies.DeviceIdGenerator.GenerateDeviceID()
	if err != nil {
		return err
	}

	err = c.dependencies.ConfigGateway.SetDeviceID(deviceID)
	if err != nil {
		return err
	}

	c.state.Status = core.DeviceStatusReady
	return nil
}
