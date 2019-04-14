package pulseaudio

import (
	"fmt"

	Config "github.com/afreakk/i3statusbear/internal/config"
	Protocol "github.com/afreakk/i3statusbear/internal/protocol"
	Util "github.com/afreakk/i3statusbear/internal/util"
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
	cl.pulse.Device(path).Get("BaseVolume", &baseVolume)
	cl.updatePulseMsg(values, baseVolume)
}

func Pulseaudio(output *Protocol.Output, module Config.Module) error {
	// === Start: PulseAudio setup ===
	pulse, e := pulseaudio.New()
	if e != nil {
		panic(e)
	}

	var pathToFallbackSink dbus.ObjectPath
	// Here we assume you are using fallbacksink, so we query that
	// altough later in pulse-callback we render whatever device you changed volume on :)
	pulse.Core().Get("FallbackSink", &pathToFallbackSink)
	volumes, e := pulse.Device(pathToFallbackSink).ListUint32("Volume")
	if e != nil {
		return e
	}
	baseVolume, e := pulse.Device(pathToFallbackSink).Uint32("BaseVolume")
	if e != nil {
		return e
	}
	formatPulseAudioText := func(volumes []uint32, baseVolume uint32) string {
		return fmt.Sprintf("vol %.2f", float32(volumes[0])/float32(baseVolume))
	}
	pulseAudioMsg := &Protocol.Message{
		FullText: formatPulseAudioText(volumes, baseVolume),
	}
	Util.ApplyModuleConfigToMessage(module, pulseAudioMsg)
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