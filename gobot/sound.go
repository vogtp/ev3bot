package gobot

import (
	"log"
	"os/exec"
	"time"

	"github.com/ev3go/ev3dev"
)

type sound struct {
	speaker *ev3dev.Speaker
}

const soundPathEv3dev = "/dev/input/by-path/platform-snd-legoev3-event"
const soundPathLego = "/dev/input/by-path/platform-sound-event"

func createSound() *sound {
	speaker := ev3dev.NewSpeaker(soundPathLego)
	if speaker == nil {
		//FIXME this is a hack
		speaker = ev3dev.NewSpeaker(soundPathEv3dev)
	}
	return &sound{speaker: speaker}
}

// Beep beeps once
func (s *sound) Beep() {
	err := s.speaker.Init()
	if err != nil {
		log.Println("Init sound ", err)
	}
	{
		defer s.speaker.Close()

		// Play tone at 440Hz for 200ms...
		err = s.speaker.Tone(600)
		if err != nil {
			log.Println("play sound sound ", err)
		}
		time.Sleep(100 * time.Millisecond)
		s.speaker.Tone(0)
	}
}

// Beep beeps n times
func (s *sound) Beeps(n int) {
	for i := 0; i < n; i++ {
		s.Beep()
		time.Sleep(100 * time.Millisecond)
	}
}

// Speak uses a goroutine to speak in the background
func (s *sound) Speak(text string) {
	go func() {
		err := exec.Command("/usr/bin/espeak", "-a 200", "-s 130", text).Run()
		if err != nil {
			log.Println("Sound.Speak: ", err)
		}
	}()
}

//espeak_opts='-a 200 -s 130'
// "/usr/bin/espeak --stdout %s '%s' | /usr/bin/aplay -q" % (espeak_opts, text)
