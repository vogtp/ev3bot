package gobot

// Brick the central block
var Brick *brick = createBrick()

// TODO add buttons

type brick struct {
	Display *display
	Sound   *sound
	LEDs    *leds
}

func createBrick() *brick {
	return &brick{
		Display: createDisplay(),
		Sound:   createSound(),
		LEDs:    createLeds(),
	}
}
