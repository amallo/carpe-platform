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

	runtime := core.NewRuntime(state, dependencies, device_routing.RouteEvent)
	runtime.Send(core.Event[any]{
		Type: core.PowerOn,
	})
	runtime.RunUntilIdle()

	// Vérifier que le deviceId est défini
	if state.DeviceID != "carpe-1234" {
		t.Errorf("State.DeviceID = %q, want %q", state.DeviceID, "carpe-1234")
	}

}

// TestBoot teste le démarrage du device
