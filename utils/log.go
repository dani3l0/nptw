package utils

import (
	"fmt"

	"github.com/fatih/color"
)

// Populated by config.go
var LogLevel int

func Log(message string) {
	log(2, " Debug ", color.FgCyan, message)
}

func Warn(message string) {
	log(1, "Warning", color.FgRed, message)
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
	fmt.Printf("[%s] %s\n", status, message)
}
