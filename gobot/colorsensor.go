package gobot

import (
	"log"
	"strconv"

	"github.com/ev3go/ev3dev"
)

// ColorSensor measures the distance
type ColorSensor struct {
	cs *ev3dev.Sensor
}

// RGB rgb color info
type RGB struct {
	r uint16
	g uint16
	b uint16
}

// NewColorSensor create an color sensor
func NewColorSensor() (*ColorSensor, error) {
	cs, err := ev3dev.SensorFor("", "lego-ev3-color")
	if err != nil {
		log.Println("us-sensor creation:", err)
	}
	sensor := &ColorSensor{cs: cs}
	return sensor, err
}

// Reflection light intensit measured by using red light
// 0 .. 100
func (s *ColorSensor) Reflection() uint8 {
	s.cs.SetMode("COL-REFLECT")
	val1, err := s.cs.Value(0)
	if err != nil {
		log.Printf("failed to read reflection: %v", err)
		return 255
	}
	val, err := strconv.ParseUint(val1, 0, 8)
	if err != nil {
		log.Printf("failed to parse reflection: %v", err)
		return 255
	}
	return uint8(val)
}

// Ambient light intensity
// 0 .. 100
func (s *ColorSensor) Ambient() uint8 {
	s.cs.SetMode("COL-AMBIENT")
	val1, err := s.cs.Value(0)
	if err != nil {
		log.Printf("failed to read ambient: %v", err)
		return 255
	}
	val, err := strconv.ParseUint(val1, 0, 8)
	if err != nil {
		log.Printf("failed to parse ambient: %v", err)
		return 255
	}
	return uint8(val)
}

// RGB color
func (s *ColorSensor) RGB() RGB {
	s.cs.SetMode("RGB-RAW")
	var err error
	var val string
	var r, g, b uint64
	val, err = s.cs.Value(0)
	if err != nil {
		log.Printf("failed to read red: %v", err)
	}
	r, err = strconv.ParseUint(val, 0, 64)
	if err != nil {
		log.Printf("failed to parse red: %v", err)
		r = 1024
	}
	val, err = s.cs.Value(1)
	if err != nil {
		log.Printf("failed to read green: %v", err)
	}
	g, err = strconv.ParseUint(val, 0, 64)
	if err != nil {
		log.Printf("failed to parse green: %v", err)
		g = 1024
	}
	val, err = s.cs.Value(2)
	if err != nil {
		log.Printf("failed to read blue: %v", err)
	}
	b, err = strconv.ParseUint(val, 0, 64)
	if err != nil {
		log.Printf("failed to parse blue: %v", err)
		b = 1024
	}
	return RGB{r: uint16(r), g: uint16(g), b: uint16(b)}
}
