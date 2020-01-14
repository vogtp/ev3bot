package gobot

import (
	"log"
	"time"

	"github.com/ev3go/ev3dev"
)

const (
	largeMotorName  = "lego-ev3-l-motor"
	mediumMotorName = "lego-ev3-m-motor"
)

// Motor is an abstraction for any motor
type Motor struct {
	TachoMotor *ev3dev.TachoMotor
}

// NewMotor is a temporary func to get an Motor
func NewMotor(port MotorPort) *Motor {
	m, err := ev3dev.TachoMotorFor(string(port), largeMotorName)
	if err != nil {
		// FIXME thats a hack
		m, err = ev3dev.TachoMotorFor(string(port), mediumMotorName)
		if err != nil {
			log.Printf("failed to find motor : %v", err)
		}
	}
	err = m.
		SetRampUpSetpoint(200 * time.Millisecond).
		SetRampDownSetpoint(200 * time.Millisecond).
		Err()
	if err != nil {
		log.Fatalf("failed to set  acceleration: %v", err)
	}
	log.Println("Motor", m, "port:", port, "maxspeed:", m.MaxSpeed(), "driver:", m.Driver(), "commands:", m.Commands(), "stopactions:", m.StopActions())
	motor := &Motor{TachoMotor: m}
	motor.SetStopAction(Stop.Coast)
	return motor
}

// SetSpeed sets the speed of the Motor
func (m *Motor) SetSpeed(speed int) error {
	if speed > m.TachoMotor.MaxSpeed() {
		speed = m.TachoMotor.MaxSpeed()
		log.Println("Setting Motorspeed to max speed ", speed)
	}
	err := m.TachoMotor.SetSpeedSetpoint(speed).Err()
	if err != nil {
		log.Printf("Motor.Setspeed: %s", err)
	}
	return err
}

// On truns the motor on
func (m *Motor) On(speed int) error {
	m.SetSpeed(speed)
	err := m.TachoMotor.Command("run-forever").Err()
	if err != nil {
		log.Printf("Motor.On: %s", err)
	}
	return err
}

// SetStopAction sets the stop action of the Motor
func (m *Motor) SetStopAction(action StopAction) *Motor {
	m.TachoMotor.SetStopAction(string(action))
	return m
}

// Stop stops the motor
func (m *Motor) Stop() error {
	return m.TachoMotor.Command("stop").Err()
}

// Err returns the error of the underling motor
func (m *Motor) Err() error {
	return m.TachoMotor.Err()
}
