package asciiart

import (
	"strconv"
)

func ColorHandler(s string, color string) []string {
	colorArray := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		colorArray[i] = "\033[0m"
	}
	if color == "rainbow" {
		for i := 0; i < len(s); i++ {
			colorArray[i] = "\033[3" + strconv.Itoa(i%7+1) + "m"
		}
	} else {
		for i := 0; i < len(s); i++ {
			colorArray[i] = "\033[3" + color + "m"
		}
	}
	return colorArray
}
