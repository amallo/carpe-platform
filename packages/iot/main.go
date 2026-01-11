package main

import (
	"machine"
	"time"
)

func main() {
	// Initialiser la LED intégrée (GPIO2 sur ESP32)
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Initialiser le port série pour le debug
	machine.Serial.Configure(machine.UARTConfig{
		BaudRate: 115200,
	})

	println("CARPE Module - ESP32 Go")
	println("Initialisation...")

	// Boucle principale
	for {
		// Allumer la LED
		led.High()
		println("LED ON")
		time.Sleep(time.Millisecond * 500)

		// Éteindre la LED
		led.Low()
		println("LED OFF")
		time.Sleep(time.Millisecond * 500)
	}
}

