# beep
A small program that plays a sound when clicking an icon in the system tray / tray bar. It is used to wake up my studio monitor speakers from the sleep mode.

Currently rewriting this in Go as a learning project.

## Notes 

Library used for system tray icon: https://github.com/getlantern/systray

To avoid opening a console at application start, use this compile flag (for Windows):

```bash
go build -ldflags -H=windowsgui
```