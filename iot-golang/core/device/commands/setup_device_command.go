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
	deviceID, err := c.dependencies.DeviceIdGenerator.GenerateDeviceID()
	if err != nil {
		return err
	}

	c.state.DeviceID = deviceID
	return nil
}
