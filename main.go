package main

import (
	"ev3gobot/gobot"
	"log"
	"os"
	"time"
)

func main() {
	// redirect log to stderr so that it prints messages in VS Code
	log.SetOutput(os.Stderr)
	var version = 3
	log.Println("Active Version", version)
	var brick = gobot.Brick
	brick.Sound.Beeps(version)
	//var err error
	left := gobot.NewMotor(gobot.Port.D)
	right := gobot.NewMotor(gobot.Port.A)
	sensorMotor := gobot.NewMotor(gobot.Port.B)
	sensorMotor.On(30)
	go func() {

		time.Sleep(time.Second * 3)
		sensorMotor.Stop()
	}()
	bot := gobot.NewGoBot(left, right)
	bot.Stop()
	// bot.Drive(100, 30)

	brick.Sound.Speak("Hello")
	brick.LEDs.GreenLeft.Off()
	brick.LEDs.GreenRight.Off()
	brick.LEDs.RedLeft.On()
	brick.LEDs.RedRight.SetBrightness(.1)
	brick.LEDs.RedRight.On()
	brick.Display.Println("Hello brick!")

	//sensor, err := gobot.NewUltrasonicSensor()
	sensor, err := gobot.NewColorSensor()
	if err != nil {
		log.Fatal(err)
	}
	for {
		val := sensor.RGB()
		log.Println(val)
		time.Sleep(200 * time.Millisecond)
	}

	

	// give some time to look at the screen before the program exits
	time.Sleep(time.Second * 5)
	brick.Sound.Beeps(3)
}
