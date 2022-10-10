package asciiart

import (
	"strconv"
)

const (
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

func ColorHandler(s string, color string) []string {
	colorArray := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		colorArray[i] = "\033[0m"
	}
	if color == "rainbow" {
		for i := 0; i < len(s); i++ {
			colorArray[i] = "\033[3" + strconv.Itoa(i%6+1) + "m"
		}
	} else {
		for i := 0; i < len(s); i++ {
			colorArray[i] = "\033[3" + color + "m"
		}
	}
	return colorArray
}
