package main

import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
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
			playBeep() // Play sound when "Play Beep" button is clicked
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

func playBeep() {
	fmt.Println("Playing beep sound!")
}

func onExit() {
}
