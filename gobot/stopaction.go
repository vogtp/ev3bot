package gobot

// StopAction is a stop action
type StopAction string

// Stop collects all available stop action
var Stop = stop{Coast: "coast", Brake: "brake", Hold: "hold"}

type stop struct {
	// StopActionCoast stops the motor without further force
	Coast StopAction
	// StopActionBrake stops the motor and resists further force
	Brake StopAction
	// StopActionHold stops the motor and activly tries to hold it
	Hold StopAction
}
