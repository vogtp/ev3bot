package gobot

import "log"

// GoBot is a drivemodule
type GoBot struct {
	leftMotor  *Motor
	rightMotor *Motor
}

// NewGoBot creates a new GoBot
func NewGoBot(leftMotor *Motor, rightMotor *Motor) *GoBot {
	gobot := &GoBot{leftMotor: leftMotor, rightMotor: rightMotor}
	return gobot
}

// Drive moves the bot
func (g *GoBot) Drive(speed int, stear int) error {

	speedLeft := speed + stear
	speedRigt := speed - stear

	g.leftMotor.On(speedLeft)
	g.rightMotor.On(speedRigt)

	err := g.leftMotor.Err()
	if err != nil {
		return err
	}
	err = g.rightMotor.Err()
	if err != nil {
		return err
	}
	return nil
}

// Stop the bot
func (g *GoBot) Stop() {
	err := g.leftMotor.Stop()
	if err != nil {
		log.Println("Stop left motor:", err)
	}
	err = g.rightMotor.Stop()
	if err != nil {
		log.Println("Stop right motor:", err)
	}
}
