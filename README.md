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
+ some go packages (go.mod)

## Install
```
go get github.com/afreakk/godwmstatus
```

## usage
```
$GOPATH/bin/godwmstatus $GOPATH/src/github.com/afreakk/godwmstatus/exampleConfigs/bar_example1.json
```
(if no GOPATH is set, it should be in ~/go)
it will log to stderr if theres any issues

Edit bar config json to match your needs.
Look at examples to understand how it works.
