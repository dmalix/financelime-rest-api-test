package utils

import "fmt"

type Color string

const (
	ColorBlack     Color = "\u001b[30m"
	ColorRed             = "\u001b[31m"
	ColorGreen           = "\u001b[32m"
	ColorYellow          = "\u001b[33m"
	ColorBlue            = "\u001b[34m"
	ColorDarkRed         = "\u001b[35m"
	ColorDarkGreen       = "\u001b[36m"
	ColorReset           = "\u001b[0m"
)

func Colorize(color Color, message string, br bool) {
	if br == true {
		fmt.Println(string(color), message, string(ColorReset))
	} else {
		fmt.Print(string(color), message, string(ColorReset))
	}
}
