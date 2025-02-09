package log

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Populated by config.go
var LogLevel int
var LogLevels = []string{"error", "warn", "info", "verbose"}

func V(message ...string) {
	log(3, "Verbose", color.FgHiMagenta, strings.Join(message, ""))
}

func I(message ...string) {
	log(2, " Info  ", color.FgCyan, strings.Join(message, ""))
}

func W(message ...string) {
	log(1, "Warning", color.FgYellow, strings.Join(message, ""))
}

func E(message ...string) {
	log(0, " Error ", color.FgRed, strings.Join(message, ""))
}

func log(level int, status string, colour color.Attribute, message string) {
	if level > LogLevel {
		return
	}
	colorFunc := color.New(colour).SprintFunc()
	status = colorFunc(status)
	now := time.Now().Format("02 Jan 15:04:05")
	for _, v := range strings.Split(message, "\n") {
		fmt.Printf("%s [%s] %s\n", now, status, v)
	}
}
