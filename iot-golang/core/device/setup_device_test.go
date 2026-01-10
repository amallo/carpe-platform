//go:build !tinygo

package device

import (
	"testing"

	"github.com/carpe-platform/iot-golang/core"
	device_gateways "github.com/carpe-platform/iot-golang/core/device/gateways/infra"
	device_generators "github.com/carpe-platform/iot-golang/core/device/generators/infra"
	device_routing "github.com/carpe-platform/iot-golang/core/device/routing"
)

// TestSetupDeviceFirstTime teste la configuration initiale du device
func TestSetupDeviceFirstTime(t *testing.T) {
	state := core.NewState()
	configGateway := device_gateways.NewFakeConfigGateway()
	deviceIdGenerator := device_generators.NewFakeDeviceIdGenerator()
	dependencies := &core.Dependencies{
		ConfigGateway:     configGateway,
		DeviceIdGenerator: deviceIdGenerator,
	}
	deviceIdGenerator.WillGenerateDeviceID("carpe-1234")

	// Vérifier le status initial
	if state.Status != core.DeviceStatusInitializing {
		t.Errorf("State.Status = %q, want %q", state.Status, core.DeviceStatusInitializing)
	}

	runtime := core.NewRuntime(state, dependencies, device_routing.RouteEvent)
	runtime.Send(core.Event[any]{
		Type: core.PowerOn,
	})
	runtime.RunUntilIdle()

	// Verify that deviceID is saved in ConfigGateway
	savedDeviceID, err := configGateway.GetDeviceID()
	if err != nil {
		t.Fatalf("Failed to get deviceID from ConfigGateway: %v", err)
	}
	if savedDeviceID != "carpe-1234" {
		t.Errorf("ConfigGateway.DeviceID = %q, want %q", savedDeviceID, "carpe-1234")
	}

	// Vérifier que le status est passé à ready
	if state.Status != core.DeviceStatusReady {
		t.Errorf("State.Status = %q, want %q", state.Status, core.DeviceStatusReady)
	}

}

// TestSetupDeviceThatwasAlreadySetup tests the setup of a device that was already configured
func TestSetupDeviceThatwasAlreadySetup(t *testing.T) {
	state := core.NewState()
	configGateway := device_gateways.NewFakeConfigGateway()
	deviceIdGenerator := device_generators.NewFakeDeviceIdGenerator()

	// Pre-configure deviceID in ConfigGateway
	existingDeviceID := "carpe-existing-1234"
	err := configGateway.SetDeviceID(existingDeviceID)
	if err != nil {
		t.Fatalf("Failed to setup existing deviceID: %v", err)
	}

	dependencies := &core.Dependencies{
		ConfigGateway:     configGateway,
		DeviceIdGenerator: deviceIdGenerator,
	}

	runtime := core.NewRuntime(state, dependencies, device_routing.RouteEvent)
	runtime.Send(core.Event[any]{
		Type: core.PowerOn,
	})
	runtime.RunUntilIdle()

	// Verify that deviceID is still in ConfigGateway
	savedDeviceID, err := configGateway.GetDeviceID()
	if err != nil {
		t.Fatalf("Failed to get deviceID from ConfigGateway: %v", err)
	}
	if savedDeviceID != existingDeviceID {
		t.Errorf("ConfigGateway.DeviceID = %q, want %q", savedDeviceID, existingDeviceID)
	}

	// Verify that status is ready
	if state.Status != core.DeviceStatusReady {
		t.Errorf("State.Status = %q, want %q", state.Status, core.DeviceStatusReady)
	}
}

// TestBoot teste le démarrage du device
