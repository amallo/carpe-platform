//go:build !tinygo

package device

import (
	"testing"

	"github.com/carpe-platform/iot-golang/core/device/test"
)

// TestSetupDeviceFirstTime teste la configuration initiale du device
func TestSetupDeviceFirstTime(t *testing.T) {
	helper := test.NewTestHelper().
		WithDeviceID("carpe-1234")

	// Given: État initial et configuration
	helper.Given(t).
		ThenAssertInitialStatus(t)

	// When: Action déclenchante
	helper.WhenPowerOn(t)

	// Then: Vérifications
	helper.
		ThenAssertDeviceID(t).
		ThenAssertStatus(t)
}

// TestSetupDeviceThatwasAlreadySetup tests the setup of a device that was already configured
func TestSetupDeviceThatwasAlreadySetup(t *testing.T) {
	existingDeviceID := "carpe-existing-1234"
	helper := test.NewTestHelper().
		WithExistingDeviceID(existingDeviceID)

	// Given: État initial avec device déjà configuré
	helper.GivenWithExistingDevice(t)

	// When: Action déclenchante
	helper.WhenPowerOn(t)

	// Then: Vérifications
	helper.
		ThenAssertDeviceID(t).
		ThenAssertStatus(t)
}

// TestBoot teste le démarrage du device
