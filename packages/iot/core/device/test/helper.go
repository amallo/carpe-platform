//go:build !tinygo

package test

import (
	"testing"

	"github.com/carpe-platform/iot-golang/core"
	device_gateways "github.com/carpe-platform/iot-golang/core/device/gateways/infra"
	device_generators "github.com/carpe-platform/iot-golang/core/device/generators/infra"
	device_routing "github.com/carpe-platform/iot-golang/core/device/routing"
)

type TestHelper struct {
	state             *core.State
	configGateway     *device_gateways.FakeConfigGateway
	deviceIdGenerator *device_generators.FakeDeviceIdGenerator
	dependencies      *core.Dependencies
	runtime           *core.Runtime

	deviceID         string
	existingDeviceID string
	expectedStatus   core.DeviceStatus
}

func NewTestHelper() *TestHelper {
	return &TestHelper{
		expectedStatus: core.DeviceStatusReady,
	}
}

func (h *TestHelper) WithDeviceID(deviceID string) *TestHelper {
	h.deviceID = deviceID
	return h
}

func (h *TestHelper) WithExistingDeviceID(deviceID string) *TestHelper {
	h.existingDeviceID = deviceID
	return h
}

func (h *TestHelper) WithExpectedStatus(status core.DeviceStatus) *TestHelper {
	h.expectedStatus = status
	return h
}

func (h *TestHelper) Given(t *testing.T) *TestHelper {
	t.Helper()
	h.state = core.NewState()
	h.configGateway = device_gateways.NewFakeConfigGateway()
	h.deviceIdGenerator = device_generators.NewFakeDeviceIdGenerator()
	h.dependencies = &core.Dependencies{
		ConfigGateway:     h.configGateway,
		DeviceIdGenerator: h.deviceIdGenerator,
	}
	if h.deviceID != "" {
		h.deviceIdGenerator.WillGenerateDeviceID(h.deviceID)
	}
	return h
}

func (h *TestHelper) GivenWithExistingDevice(t *testing.T) *TestHelper {
	t.Helper()
	h.Given(t)
	if h.existingDeviceID != "" {
		err := h.configGateway.SetDeviceID(h.existingDeviceID)
		if err != nil {
			t.Fatalf("Failed to setup existing deviceID: %v", err)
		}
	}
	return h
}

func (h *TestHelper) WhenPowerOn(t *testing.T) *TestHelper {
	t.Helper()
	h.runtime = core.NewRuntime(h.state, h.dependencies, device_routing.RouteEvent)
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

func (h *TestHelper) ThenAssertDeviceID(t *testing.T) *TestHelper {
	t.Helper()
	expectedID := h.deviceID
	if h.existingDeviceID != "" {
		expectedID = h.existingDeviceID
	}
	savedDeviceID, err := h.configGateway.GetDeviceID()
	if err != nil {
		t.Fatalf("Failed to get deviceID from ConfigGateway: %v", err)
	}
	if savedDeviceID != expectedID {
		t.Errorf("ConfigGateway.DeviceID = %q, want %q", savedDeviceID, expectedID)
	}
	return h
}

func (h *TestHelper) ThenAssertStatus(t *testing.T) *TestHelper {
	t.Helper()
	if h.state.Status != h.expectedStatus {
		t.Errorf("State.Status = %q, want %q", h.state.Status, h.expectedStatus)
	}
	return h
}
