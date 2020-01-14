package gobot

// MotorPort is the porttype for motors
type MotorPort string

// Port is a collection of possible ports
var Port port = port{
	A: "ev3-ports:outA",
	B: "ev3-ports:outB",
	C: "ev3-ports:outC",
	D: "ev3-ports:outD"}

type port struct {
	// A motor port
	A MotorPort
	// B motor port
	B MotorPort
	// C motor port
	C MotorPort
	// D motor port
	D MotorPort
}
