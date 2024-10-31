package main

import (
	"embed"
	"os"
	"sync"
	"time"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	setup()
}

func setup() {
	systray.SetTitle("Beep")
	systray.SetTooltip("Play a sound to wake up the studio monitors")
	systray.SetIcon(icon.Data) // TODO - change this to a custom icon

	beepButton := systray.AddMenuItem("Play Beep", "Play a beep sound")
	quitButton := systray.AddMenuItem("Quit", "Exit the application")

	go func() {
		for {
			<-beepButton.ClickedCh
			playSound()
		}
	}()

	go func() {
		for {
			<-quitButton.ClickedCh
			systray.Quit()
			os.Exit(0)
		}
	}()
}

//go:embed beep.wav
var soundFile embed.FS
var speakerInitialized bool

func playSound() {
	if !speakerInitialized {
		sampleRate := beep.SampleRate(48000)
		if err := speaker.Init(sampleRate, sampleRate.N(time.Second/10)); err != nil {
			panic(err)
		}
		speakerInitialized = true
	}

	// Open the embedded sound file
	f, err := soundFile.Open("beep.wav")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Decode the WAV file
	decoder, _, err := wav.Decode(f)
	if err != nil {
		panic(err)
	}
	defer decoder.Close()

	// wait group to wait for the sound to finish
	var wg sync.WaitGroup
	wg.Add(1)

	speaker.Play(beep.Seq(decoder, beep.Callback(func() {
		wg.Done() // playback done
	})))

	// Wait for the sound to finish
	wg.Wait()
}

func onExit() {
	if speakerInitialized {
		speaker.Close()
	}
}
