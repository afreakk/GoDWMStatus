package protocol

// #cgo LDFLAGS: -lX11
// #include <X11/Xlib.h>
import "C"

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/afreakk/godwmstatus/internal/config"
)

var dpy = C.XOpenDisplay(nil)

type Output struct {
	Messages             []*Message
	renderTimer          *time.Timer
	renderTimerIsRunning bool
	renderInterval       time.Duration
	mux                  sync.Mutex
}

func (o *Output) Init(cfg config.Config) (err error) {
	o.Messages = []*Message{}
	if o.renderInterval, err = time.ParseDuration(cfg.MinimumRenderInterval); err != nil {
		return
	}
	o.renderTimer = time.AfterFunc(o.renderInterval, o.ActuallyPrintMsgs)
	return
}

func (o *Output) PrintMsgs() {
	o.mux.Lock()
	if !o.renderTimerIsRunning {
		o.renderTimer.Reset(o.renderInterval)
		o.renderTimerIsRunning = true
	}
	o.mux.Unlock()
}

func printToStderrIfErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func (o *Output) ActuallyPrintMsgs() {
	o.renderTimerIsRunning = false
	var b strings.Builder
	for _, msg := range o.Messages {
		b.WriteString(msg.FullText)
	}
	C.XStoreName(dpy, C.XDefaultRootWindow(dpy), C.CString(b.String()))
	C.XSync(dpy, 1)
}

type Message struct {
	FullText string `json:"full_text"`
}
