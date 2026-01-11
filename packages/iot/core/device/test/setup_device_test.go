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

// TestSetupDeviceWhenDeviceIdGenerationFails teste le cas où la génération du device ID échoue
func TestSetupDeviceWhenDeviceIdGenerationFails(t *testing.T) {
	helper := NewTestHelper()

	// Given: État initial avec générateur configuré pour échouer
	helper.GivenWillFailToGenerateDeviceId(t, "failed to generate device ID").
		ThenAssertInitialStatus(t)

	// When: Action déclenchante
	helper.WhenPowerOn(t)

	// Then: Vérifications
	helper.
		ThenAssertStatus(t, core.DeviceStatusErrorCannotGenerateDeviceId)

	// Vérifier que le device ID n'a pas été configuré
	helper.ThenAssertDeviceIDWasConfigured(t, "")
}

// TestBoot teste le démarrage du device
