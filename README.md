# GoDWMStatus
### Works with DWM statusbar

## Modules
- command (run bash scripts etc)
- cpu (shows usage)
- datetime
- memory (shows usage)
- pulseaudio (volume bar, eventdriven, not polling)
- readfile (read arbitrary file)

## Dependency
X11/Xlib.h
And some go packages `go get ./..`

## Example usage
```
godwmstatus /path/to/bar_example1.json 2>> ~/log/godwmstatus.log
```
