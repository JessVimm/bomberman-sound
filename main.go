package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	// Open
	sound, err := os.Open("Bomberman.mp3")
	if err != nil {
		log.Fatal(err)
	}

	// Decode
	streamer, format, err := mp3.Decode(sound)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	// Initialize speaker
	sr := format.SampleRate * 2
	speaker.Init(sr, sr.N(time.Second/10))
	resampled := beep.Resample(4, format.SampleRate, sr, streamer)

	// Play
	stopEntry := ""
	speaker.Play(beep.Seq(resampled, beep.Callback(func() {})))
	
	// Stop sound
	fmt.Print("Stop? -> ")
	fmt.Scanln(&stopEntry)


	// Clear
	speaker.Clear()
	
	fmt.Println("We are done!")
	
}
