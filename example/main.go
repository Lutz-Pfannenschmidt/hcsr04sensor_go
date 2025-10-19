package main

import (
	"time"

	hcsr04 "github.com/Lutz-Pfannenschmidt/hcsr04sensor_go"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	trigPin := rpio.Pin(23)
	echoPin := rpio.Pin(24)

	// !!! DO NOT USE !!!
	// if you have multiple sensors or other gpio in your program,
	// call rpio.Open() once in your main program instead of in InitPins.
	// and initialize the pins manually.
	hcsr04.InitPins(trigPin, echoPin)

	distance, err := hcsr04.Measure(trigPin, echoPin, 20.0, 11, 100*time.Millisecond)
	if err != nil {
		panic(err)
	}
	println("Distance (mm):", distance)

	// Again, only call if you don't use gpio elsewhere in your program.
	hcsr04.ClosePins()
}
