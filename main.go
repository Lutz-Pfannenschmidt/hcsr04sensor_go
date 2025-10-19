package hcsr04

import (
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

// InitPins initializes the trigger and echo pins for the HC-SR04 sensor.
// If using more than one sensor, consider calling rpio.Open() and setting the pins up once in your main program instead.
func InitPins(trig, echo rpio.Pin) error {
	if err := rpio.Open(); err != nil {
		return fmt.Errorf("failed to open rpio: %v", err)
	}
	trig.Output()
	echo.Input()
	echo.PullOff()
	return nil
}

// ClosePins closes the rpio library.
// If using gpio in other parts of your program, do not call this function until the program is ending.
func ClosePins() {
	rpio.Close()
}

// Measure measures the distance using the HC-SR04 sensor.
// It returns the median distance in mm based on the specified number of samples.
func Measure(trig, echo rpio.Pin, temperatureC float64, samples int, wait time.Duration) (int, error) {
	speedOfSound := 331.3 * math.Sqrt(1+(temperatureC/273.15)) // m/s

	var readings []int
	for i := 0; i < samples; i++ {
		trig.Low()
		time.Sleep(wait)

		trig.High()
		time.Sleep(10 * time.Microsecond)
		trig.Low()

		start := time.Now()
		timeout := start.Add(1 * time.Second)

		for echo.Read() == rpio.Low {
			if time.Now().After(timeout) {
				return 0, fmt.Errorf("echo pulse start timeout")
			}
		}
		signalStart := time.Now()

		for echo.Read() == rpio.High {
			if time.Now().After(timeout) {
				return 0, fmt.Errorf("echo pulse end timeout")
			}
		}
		signalEnd := time.Now()

		duration := signalEnd.Sub(signalStart).Seconds()
		distance := int(duration*speedOfSound*1000) / 2

		readings = append(readings, distance)
		time.Sleep(wait)
	}

	sort.Ints(readings)
	return readings[len(readings)/2], nil // median
}
