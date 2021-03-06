package pulseaudio

import (
	"fmt"
	"os"

	"github.com/afreakk/godwmstatus/internal/config"
	"github.com/afreakk/godwmstatus/internal/protocol"
	"github.com/afreakk/godwmstatus/internal/util"
	"github.com/godbus/dbus"
	"github.com/sqp/pulseaudio"
)

// Pulse audio subscribe stuff
type Client struct {
	pulse          *pulseaudio.Client
	updatePulseMsg func(volumes []uint32, baseVolume uint32)
}

func (cl *Client) DeviceVolumeUpdated(path dbus.ObjectPath, values []uint32) {
	var baseVolume uint32
	err := cl.pulse.Device(path).Get("BaseVolume", &baseVolume)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	cl.updatePulseMsg(values, baseVolume)
}

func loadPulseAudioModuleIfNeeded() error {
	isLoaded, err := pulseaudio.ModuleIsLoaded()
	if err != nil {
		return err
	}
	if !isLoaded {
		err = pulseaudio.LoadModule()
		if err != nil {
			return err
		}
	}
	return nil
}

func Pulseaudio(output *protocol.Output, module config.Module) error {
	if errLoadingModule := loadPulseAudioModuleIfNeeded(); errLoadingModule != nil {
		return errLoadingModule
	}
	// === Start: PulseAudio setup ===
	pulse, e := pulseaudio.New()
	if e != nil {
		return e
	}

	var pathToFallbackSink dbus.ObjectPath
	// Here we assume you are using fallbacksink, so we query that
	// altough later in pulse-callback we render whatever device you changed volume on :)
	getFallbackSinkErr := pulse.Core().Get("FallbackSink", &pathToFallbackSink)
	if getFallbackSinkErr != nil {
		return getFallbackSinkErr
	}
	volumes, e := pulse.Device(pathToFallbackSink).ListUint32("Volume")
	if e != nil {
		return e
	}
	baseVolume, e := pulse.Device(pathToFallbackSink).Uint32("BaseVolume")
	if e != nil {
		return e
	}
	formatPulseAudioText := func(volumes []uint32, baseVolume uint32) string {
		return util.RenderBar(module, int64(volumes[0]), int64(baseVolume))
	}
	pulseAudioMsg := &protocol.Message{
		FullText: formatPulseAudioText(volumes, baseVolume),
	}
	updatePulseMsg := func(volumes []uint32, baseVolume uint32) {
		pulseAudioMsg.FullText = formatPulseAudioText(volumes, baseVolume)
		output.PrintMsgs()
	}

	client := &Client{pulse, updatePulseMsg}
	pulse.Register(client)

	go pulse.Listen()
	// === End: PulseAudio setup ===
	output.Messages = append(output.Messages, pulseAudioMsg)
	return nil
}
