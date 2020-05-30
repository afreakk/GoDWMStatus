package readfile

import (
	"fmt"
	"io/ioutil"

	"github.com/afreakk/godwmstatus/internal/config"
	"github.com/afreakk/godwmstatus/internal/protocol"
)

func Readfile(output *protocol.Output, module config.Module) func() {
	formatString := func() string {
		data, err := ioutil.ReadFile(module.FilePath)
		if err != nil {
			return err.Error()
		}
		return fmt.Sprintf(module.Sprintf, string(data[:len(data)-1]))
	}
	fileMsg := &protocol.Message{
		FullText: formatString(),
	}
	output.Messages = append(output.Messages, fileMsg)
	var lastFullText string
	return func() {
		fileMsg.FullText = formatString()
		if lastFullText != fileMsg.FullText {
			output.PrintMsgs()
			lastFullText = fileMsg.FullText
		}
	}
}
