package gobot

import (
	"log"
	"math"
	"strconv"

	"github.com/ev3go/ev3dev"
)

// UltrasonicSensor measures the distance
type UltrasonicSensor struct {
	us *ev3dev.Sensor
}

// NewUltrasonicSensor create an ultrasonic sensor
func NewUltrasonicSensor() (*UltrasonicSensor, error) {
	us, err := ev3dev.SensorFor("", "lego-ev3-us")
	if err != nil {
		log.Println("us-sensor creation:", err)
	}
	log.Println("US:", us.Driver(), "Path", us.Path(), us.String())
	sensor := &UltrasonicSensor{us: us}
	return sensor, err
}

// Distance get the distance
func (s *UltrasonicSensor) Distance() float64 {
	d := s.us.Decimals()
	corr := 1 / math.Pow10(d)

	ds, err := s.us.Value(0)
	if err != nil {
		log.Fatalf("failed to read distance: %v", err)
	}
	dist, err := strconv.ParseFloat(ds, 64)
	if err != nil {
		log.Fatalf("failed to parse distance: %v", err)
	}
	dist *= corr
	return dist
}
