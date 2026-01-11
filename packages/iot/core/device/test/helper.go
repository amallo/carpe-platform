//go:build !tinygo

package test

import (
	"errors"
	"testing"

	"github.com/carpe-platform/iot-golang/core"
	device_gateways "github.com/carpe-platform/iot-golang/core/device/gateways/infra"
	device_generators "github.com/carpe-platform/iot-golang/core/device/generators/infra"
	device_reducers "github.com/carpe-platform/iot-golang/core/device/reducers"
	device_routing "github.com/carpe-platform/iot-golang/core/device/routing"
)

type TestHelper struct {
	state             *core.State
	configGateway     *device_gateways.FakeConfigGateway
	deviceIdGenerator *device_generators.FakeDeviceIdGenerator
	dependencies      *core.Dependencies
	runtime           *core.Runtime
}

func NewTestHelper() *TestHelper {
	return &TestHelper{}
}

func (h *TestHelper) GivenWillGenerateDeviceId(t *testing.T, deviceID string) *TestHelper {
	t.Helper()
	h.state = core.NewState()
	h.configGateway = device_gateways.NewFakeConfigGateway()
	h.deviceIdGenerator = device_generators.NewFakeDeviceIdGenerator()
	h.dependencies = &core.Dependencies{
		ConfigGateway:     h.configGateway,
		DeviceIdGenerator: h.deviceIdGenerator,
	}
	h.deviceIdGenerator.WillGenerateDeviceID(deviceID)
	return h
}

func (h *TestHelper) GivenWithExistingDevice(t *testing.T, existingDeviceID string) *TestHelper {
	t.Helper()
	h.state = core.NewState()
	h.configGateway = device_gateways.NewFakeConfigGateway()
	h.deviceIdGenerator = device_generators.NewFakeDeviceIdGenerator()
	h.dependencies = &core.Dependencies{
		ConfigGateway:     h.configGateway,
		DeviceIdGenerator: h.deviceIdGenerator,
	}
	err := h.configGateway.SetDeviceID(existingDeviceID)
	if err != nil {
		t.Fatalf("Failed to setup existing deviceID: %v", err)
	}
	return h
}

func (h *TestHelper) GivenWillFailToGenerateDeviceId(t *testing.T, errorMessage string) *TestHelper {
	t.Helper()
	h.state = core.NewState()
	h.configGateway = device_gateways.NewFakeConfigGateway()
	h.deviceIdGenerator = device_generators.NewFakeDeviceIdGenerator()
	h.dependencies = &core.Dependencies{
		ConfigGateway:     h.configGateway,
		DeviceIdGenerator: h.deviceIdGenerator,
	}
	h.deviceIdGenerator.WillFailWithError(errors.New(errorMessage))
	return h
}

func (h *TestHelper) WhenPowerOn(t *testing.T) *TestHelper {
	t.Helper()
	reducers := []core.EventReducer{
		device_reducers.ReduceDeviceEvents,
	}
	h.runtime = core.NewRuntime(h.state, h.dependencies, device_routing.RouteEvent, reducers)
	h.runtime.Send(core.Event[any]{
		Type: core.PowerOn,
	})
	h.runtime.RunUntilIdle()
	return h
}

func (h *TestHelper) ThenAssertInitialStatus(t *testing.T) *TestHelper {
	t.Helper()
	if h.state.Status != core.DeviceStatusInitializing {
		t.Errorf("State.Status = %q, want %q", h.state.Status, core.DeviceStatusInitializing)
	}
	return h
}

func (h *TestHelper) ThenAssertDeviceIDWasConfigured(t *testing.T, expectedDeviceID string) *TestHelper {
	t.Helper()
	savedDeviceID, err := h.configGateway.GetDeviceID()
	if err != nil {
		t.Fatalf("Failed to get deviceID from ConfigGateway: %v", err)
	}
	if savedDeviceID != expectedDeviceID {
		t.Errorf("ConfigGateway.DeviceID = %q, want %q", savedDeviceID, expectedDeviceID)
	}
	return h
}

func (h *TestHelper) ThenAssertStatus(t *testing.T, expectedStatus core.DeviceStatus) *TestHelper {
	t.Helper()
	if h.state.Status != expectedStatus {
		t.Errorf("State.Status = %q, want %q", h.state.Status, expectedStatus)
	}
	return h
}
