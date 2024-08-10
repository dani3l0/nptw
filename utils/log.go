package utils

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// Populated by config.go
var LogLevel int

func Log(message string) {
	log(2, " Debug ", color.FgCyan, message)
}

func Warn(message string) {
	log(1, "Warning", color.FgYellow, message)
}

func Err(message string) {
	log(0, " Error ", color.FgRed, message)
}

func log(level int, status string, colour color.Attribute, message string) {
	if level > LogLevel {
		return
	}
	colorFunc := color.New(colour).SprintFunc()
	status = colorFunc(status)
	for _, v := range strings.Split(message, "\n") {
		fmt.Printf("[%s] %s\n", status, v)
	}
}
