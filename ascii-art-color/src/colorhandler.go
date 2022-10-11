package asciiart

import (
	"strconv"
	"strings"
)

func ColorHandler(s string, color string) []string {
	colorArray := make([]string, len(s))
	colorMap, colorChars := getColoredCharacters(color)
	println(len(colorMap))
	println(len(colorChars))
	if strings.HasPrefix(color, "rainbow") {
		for i := 0; i < len(s); i++ {
			if strings.ContainsAny(string(s[i]), colorChars) || len(colorChars) == 0 {
				for currentColor := range colorMap {
					if strings.ContainsAny(string(s[i]), colorMap[currentColor]) || len(colorChars) == 0 {
						colorArray[i] = "\033[3" + strconv.Itoa(i%6+1) + "m"
						break
					}
				}
			} else {
				colorArray[i] = "\033[0m"
			}
		}
	} else {
		for i := 0; i < len(s); i++ {
			if strings.ContainsAny(string(s[i]), colorChars) || len(colorChars) == 0 {
				for currentColor := range colorMap {
					if strings.ContainsAny(string(s[i]), colorMap[currentColor]) || len(colorChars) == 0 {
						colorArray[i] = convertColor(currentColor)
						break
					}
				}
			} else {
				colorArray[i] = "\033[0m"
			}
		}
	}
	return colorArray
}

func getColoredCharacters(color string) (map[string]string, string) {
	color = strings.ToLower(color)
	mappedChars := make(map[string]string)
	allChars := ""
	currentColor := ""
	for i := 0; i < len(color); i++ {
		if currentColor == "" {
			for color[i] != ':' && i != len(color)-1 {
				currentColor += string(color[i])
				i++
			}
			if i == len(color)-1 {
				currentColor += string(color[i])
			}
			mappedChars[currentColor] = ""
		} else {
			for color[i] != ',' && i != len(color)-1 {
				mappedChars[currentColor] += string(color[i])
				allChars += string(color[i])
				i++
			}
			if color[i] == ',' {
				currentColor = ""
			} else {
				mappedChars[currentColor] += string(color[i])
				allChars += string(color[i])
			}
		}
	}
	return mappedChars, allChars
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
