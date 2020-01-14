package gobot

import (
	"log"

	"github.com/ev3go/ev3dev"

	"github.com/ev3go/ev3"
)

type leds struct {
	RedLeft    *led
	RedRight   *led
	GreenLeft  *led
	GreenRight *led
}

func createLeds() *leds {
	leds := &leds{
		RedLeft:    &led{led: ev3.RedLeft},
		RedRight:   &led{led: ev3.RedRight},
		GreenLeft:  &led{led: ev3.GreenLeft},
		GreenRight: &led{led: ev3.GreenRight},
	}
	return leds
}

type led struct {
	led           *ev3dev.LED
	brightness    int
	maxBrightness int
}

// GetBrightness gets the brightness of the LED
func (l *led) GetBrightness() int {
	if l.brightness == 0 {
		l.brightness = l.getMaxBrightness()
	}
	return l.brightness
}

// GetBrightness sets the brightness of the LED
// without turning the LED on
// 0.0 .. 1.0
func (l *led) SetBrightness(brightness float32) int {
	if l.maxBrightness == 0 {
		l.maxBrightness = l.getMaxBrightness()
	}
	l.brightness = int(float32(l.maxBrightness) * brightness)
	return l.brightness
}

func (l *led) getMaxBrightness() int {
	if l.maxBrightness == 0 {
		var err error
		l.maxBrightness, err = l.led.MaxBrightness()
		if err != nil {
			log.Println("LEDs: get max brightness: ", err)
		}
	}
	return l.maxBrightness
}

// On switch LED on
func (l *led) On() *led {
	l.led.SetBrightness(l.GetBrightness())
	return l
}

// Off switch off
func (l *led) Off() *led {
	l.led.SetBrightness(0)
	return l
}

// Err display errors
func (l *led) Err() error {
	return l.led.Err()
}
