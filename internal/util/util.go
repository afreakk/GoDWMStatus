package util

import (
	"fmt"
	"strings"

	"github.com/afreakk/godwmstatus/internal/config"
)

func RenderBar(module config.Module, filled int64, total int64) string {
	filledBarsCount := filled * module.BarWidth / total
	remainingCount := Max(module.BarWidth-filledBarsCount, 0)
	filledBars := strings.Repeat(module.BarFilled, int(filledBarsCount))
	emptyBars := strings.Repeat(module.BarEmpty, int(remainingCount))
	return fmt.Sprintf(module.Sprintf, filledBars, emptyBars)
}
func Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
