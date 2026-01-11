//go:build !tinygo

package test

import (
	"testing"

	"github.com/carpe-platform/iot-golang/core"
)

// TestSetupDeviceFirstTime teste la configuration initiale du device
func TestSetupDeviceFirstTime(t *testing.T) {
	helper := NewTestHelper()

	// Given: État initial et configuration
	helper.GivenWillGenerateDeviceId(t, "carpe-1234").
		ThenAssertInitialStatus(t)

	// When: Action déclenchante
	helper.WhenPowerOn(t)

	// Then: Vérifications
	helper.
		ThenAssertDeviceIDWasConfigured(t, "carpe-1234").
		ThenAssertStatus(t, core.DeviceStatusReady)
}

// TestSetupDeviceThatwasAlreadySetup tests the setup of a device that was already configured
func TestSetupDeviceThatwasAlreadySetup(t *testing.T) {
	existingDeviceID := "carpe-existing-1234"
	helper := NewTestHelper()

	// Given: État initial avec device déjà configuré
	helper.GivenWithExistingDevice(t, existingDeviceID)

	// When: Action déclenchante
	helper.WhenPowerOn(t)

	// Then: Vérifications
	helper.
		ThenAssertDeviceIDWasConfigured(t, existingDeviceID).
		ThenAssertStatus(t, core.DeviceStatusReady)
}

// TestBoot teste le démarrage du device
