# beep
Plays a sound when clicking an icon in the system tray. It is used to wake up my studio monitor speakers from the sleep mode.

To avoid opening a console at application start, use this compile flag (for Windows):

```bash
go build -ldflags -H=windowsgui
```

Or download a pre-built .exe in the [releases](https://github.com/matkv/beep/releases)
