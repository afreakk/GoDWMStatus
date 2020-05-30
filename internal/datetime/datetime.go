package datetime

import (
	"fmt"
	"time"

	"github.com/afreakk/godwmstatus/internal/config"
	"github.com/afreakk/godwmstatus/internal/protocol"
)

func Datetime(output *protocol.Output, module config.Module) func() {
	formatDateTimeMsg := func() string {
		return fmt.Sprintf(module.Sprintf, time.Now().Format(module.DateTimeFormat))
	}
	dateTimeMsg := &protocol.Message{
		FullText: formatDateTimeMsg(),
	}
	output.Messages = append(output.Messages, dateTimeMsg)
	return func() {
		dateTimeMsg.FullText = formatDateTimeMsg()
		output.PrintMsgs()
	}
}
