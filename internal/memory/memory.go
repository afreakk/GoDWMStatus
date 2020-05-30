package memory

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/afreakk/godwmstatus/internal/config"
	"github.com/afreakk/godwmstatus/internal/protocol"
	"github.com/afreakk/godwmstatus/internal/util"
)

func Memory(output *protocol.Output, module config.Module) func() {
	formatString := func() string {
		f, err := os.Open("/proc/meminfo")
		if err != nil {
			return err.Error()
		}
		fScanner := bufio.NewScanner(f)
		var line string
		var memtotal int64
		var memavail int64
		for fScanner.Scan() {
			line = fScanner.Text()
			if strings.HasPrefix(line, "MemTotal") {
				memtotal, err = strconv.ParseInt(strings.Fields(line)[1], 10, 32)
				if err != nil {
					return err.Error()
				}
			} else if strings.HasPrefix(line, "MemAvailable") {
				memavail, err = strconv.ParseInt(strings.Fields(line)[1], 10, 32)
				if err != nil {
					return err.Error()
				}
			}
			if memtotal != 0 && memavail != 0 {
				break
			}
		}
		f.Close()
		return util.RenderBar(module, memtotal-memavail, memtotal)
	}
	memMsg := &protocol.Message{
		FullText: formatString(),
	}
	output.Messages = append(output.Messages, memMsg)
	var lastFullText string
	return func() {
		memMsg.FullText = formatString()
		if lastFullText != memMsg.FullText {
			output.PrintMsgs()
			lastFullText = memMsg.FullText
		}
	}
}
