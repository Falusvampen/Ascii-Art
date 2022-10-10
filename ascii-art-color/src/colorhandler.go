package asciiart

import (
	"strconv"
	"strings"
)

func ColorHandler(s string, color string) []string {
	colorArray := make([]string, len(s))
	colorChars := getColoredCharacters(color)
	if strings.HasPrefix(color, "rainbow") {
		for i := 0; i < len(s); i++ {
			if len(colorChars) == 0 {
				colorArray[i] = "\033[3" + strconv.Itoa(i%6+1) + "m"
			} else {
				for colors := range colorChars {
					colorArray[i] = "\033[0m"
					if strings.ContainsAny(string(s[i]), colorChars[colors]) {
						colorArray[i] = "\033[3" + strconv.Itoa(i%6+1) + "m"
					}
				}
			}
		}
	} else {
		for i := 0; i < len(s); i++ {
			colorArray[i] = "\033[0m"
			for currentColor := range colorChars {
				if strings.ContainsAny(string(s[i]), colorChars[currentColor]) {
					colorArray[i] = convertColor(currentColor)
					break
				}
			}
		}
	}
	return colorArray
}

func getColoredCharacters(color string) map[string]string {
	color = strings.ToLower(color)
	chars := make(map[string]string)
	currentColor := ""
	for i := 0; i < len(color); i++ {
		if color[i] == ',' {
			currentColor = ""
			i++
		} else if currentColor == "" {
			for color[i] != ':' && i != len(color)-1 {
				currentColor += string(color[i])
				i++
			}
		} else {
			for color[i] != ',' && i != len(color)-1 {
				chars[currentColor] += string(color[i])
				i++
			}
			chars[currentColor] += string(color[i])
			break
		}
	}
	return chars
}

func convertColor(color string) string {
	switch strings.ToLower(color) {
	case "black":
		return "\033[30m"
	case "red":
		return "\033[31m"
	case "green":
		return "\033[32m"
	case "yellow":
		return "\033[33m"
	case "blue":
		return "\033[34m"
	case "magenta":
		return "\033[35m"
	case "cyan":
		return "\033[36m"
	case "white":
		return "\033[37m"
	default:
		return "\033[0m"
	}
}
