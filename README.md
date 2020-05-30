# GoDWMStatus
## Sets DWM status text

## Modules
- command (run bash scripts etc)
- cpu (usage)
- datetime
- memory (usage)
- pulseaudio (volume bar)
- readfile (read arbitrary file)

## Dependency
X11/Xlib.h
And some go packages `go get ./..`

## Example usage
```
~/go/bin/godwmstatus ~/go/src/github.com/afreakk/godwmstatus/exampleConfigs/mainbar.json i3 2>> ~/log/godwmstatus.log
```
