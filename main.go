package main

import (
	"embed"
	"os"
	"sync"
	"time"

	"github.com/getlantern/systray"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

func main() {
	systray.Run(onReady, onExit)
}

//go:embed beep.wav
var soundFile embed.FS

//go:embed icons/play.ico
var playIcon []byte

//go:embed icons/close.ico
var closeIcon []byte

func onReady() {
	setup()
}

func setup() {
	systray.SetTitle("Beep")
	systray.SetTooltip("Play a sound to wake up the studio monitors")
	systray.SetIcon(playIcon)

	beepButton := systray.AddMenuItem("Play Beep", "Play a beep sound")
	beepButton.SetIcon(playIcon)

	closeButton := systray.AddMenuItem("Quit", "Exit the application")
	closeButton.SetIcon(closeIcon)

	go func() {
		for {
			<-beepButton.ClickedCh
			playSound()
		}
	}()

	go func() {
		for {
			<-closeButton.ClickedCh
			systray.Quit()
			os.Exit(0)
		}
	}()
}

var speakerInitialized bool

func playSound() {
	if !speakerInitialized {
		sampleRate := beep.SampleRate(48000)
		if err := speaker.Init(sampleRate, sampleRate.N(time.Second/10)); err != nil {
			panic(err)
		}
		speakerInitialized = true
	}

	f, err := soundFile.Open("beep.wav")
	if err != nil {
		panic(err)
	}
	defer f.Close()

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

	// wait for sound to finish
	wg.Wait()
}

func onExit() {
	if speakerInitialized {
		speaker.Close()
	}
}
