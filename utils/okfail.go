package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func Check(message string) {
	message += "..."
	fmt.Printf("%-*s", 40, message)
}

func OkFail(result bool) {
	var msg string
	red := color.New(color.FgRed).SprintFunc()
	grn := color.New(color.FgGreen).SprintFunc()
	if result {
		msg = grn(" OK ")
	} else {
		msg = red("Fail")
	}
	fmt.Printf("[%s]\n", msg)
}
