package command

import (
	"fmt"
	"os/exec"

	"github.com/afreakk/godwmstatus/internal/config"
	"github.com/afreakk/godwmstatus/internal/protocol"
)

func Command(output *protocol.Output, module config.Module) func() {
	formatString := func() string {
		cmd := exec.Command(module.CommandName, module.CommandArgs...)
		out, err := cmd.Output()
		if err != nil {
			return err.Error()
		}

		return fmt.Sprintf(module.Sprintf, string(out[:len(out)-1]))
	}
	cmdMsg := &protocol.Message{
		FullText: formatString(),
	}
	output.Messages = append(output.Messages, cmdMsg)
	var lastFullText string
	return func() {
		cmdMsg.FullText = formatString()
		if lastFullText != cmdMsg.FullText {
			output.PrintMsgs()
			lastFullText = cmdMsg.FullText
		}
	}
}
