# hc-sr04-go

Go (Golang) implementation of the HC‑SR04 Ultrasonic Sensor measurement routines  
for Raspberry Pi using the [`go-rpio`](https://github.com/stianeikeland/go-rpio) GPIO library.

This project is a **metric‑only port** of the Python package  
[`hcsr04sensor`](https://pypi.org/project/hcsr04sensor/) by Al Audet.

---

## Overview

The HC‑SR04 ultrasonic sensor measures distance by emitting a short ultrasonic pulse
and timing its echo reflection.  
This Go version replicates the core logic from the Python `Measurement` class:

-   Median distance sampling for error correction
-   Temperature compensation for the speed of sound
-   Safe GPIO handling using [`go-rpio`](https://github.com/stianeikeland/go-rpio)

All readings are in **millimeters** only.  
Imperial conversions and unrelated API methods have been removed.

---

## Usage

```bash
go get github.com/Lutz-Pfannenschmidt/hcsr04sensor_go
```

```go
package main

import (
	"time"

	hcsr04 "github.com/Lutz-Pfannenschmidt/hcsr04sensor_go"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	trigPin := rpio.Pin(23)
	echoPin := rpio.Pin(24)

	// Do not use if you have multiple sensors or other gpio in your program,
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
```
