package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/afreakk/godwmstatus/internal/command"
	"github.com/afreakk/godwmstatus/internal/config"
	"github.com/afreakk/godwmstatus/internal/cpu"
	"github.com/afreakk/godwmstatus/internal/datetime"
	"github.com/afreakk/godwmstatus/internal/memory"
	"github.com/afreakk/godwmstatus/internal/protocol"
	"github.com/afreakk/godwmstatus/internal/pulseaudio"
	"github.com/afreakk/godwmstatus/internal/readfile"
	"github.com/robfig/cron"
)

func errorExit(str string, code int) {
	fmt.Fprintln(os.Stderr, str)
	os.Exit(code)
}

func main() {
	osArgsLen := len(os.Args)

	if osArgsLen < 2 {
		errorExit("Please provide config as argument", 1)
	}
	configFilePath := os.Args[1]

	cfg, err := config.GetConfigFromPath(configFilePath)
	if err != nil {
		panic(err)
	}

	output := protocol.Output{}
	err = output.Init(cfg)
	if err != nil {
		panic(err)
	}

	c := cron.New()
	for _, module := range cfg.Modules {
		var err error
		switch module.Name {
		case "datetime":
			_, err = c.AddFunc(module.Cron, datetime.Datetime(&output, module))
		case "pulseaudio":
			err = pulseaudio.Pulseaudio(&output, module)
		case "readfile":
			_, err = c.AddFunc(module.Cron, readfile.Readfile(&output, module))
		case "memory":
			_, err = c.AddFunc(module.Cron, memory.Memory(&output, module))
		case "cpu":
			_, err = c.AddFunc(module.Cron, cpu.Cpu(&output, module))
		case "command":
			_, err = c.AddFunc(module.Cron, command.Command(&output, module))
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "\nError when initializing module: %s, \nerror: %s\n", module.Name, err.Error())
		}
	}
	output.PrintMsgs()
	c.Start()
	runtime.Goexit()
}
