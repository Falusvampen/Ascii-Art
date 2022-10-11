package asciiart

import (
	"strconv"
	"strings"
)

func ColorHandler(s string, color string) []string {
	colorArray := make([]string, len(s))
	colorMap, colorChars := getColoredCharacters(color)
	if colorMap == nil {
		return nil
	}
	colorArray = setColoredCharacters(s, color, colorMap, colorChars, colorArray)
	return colorArray
}

func setColoredCharacters(s string, color string, colorMap map[string]string, colorChars string, colorArray []string) []string {
	for i := 0; i < len(s); i++ {
		if strings.ContainsAny(string(s[i]), colorChars) || strings.ContainsAny(strconv.Itoa(i), colorChars) || len(colorChars) == 0 {
			for currentColor := range colorMap {
				if strings.ContainsAny(string(s[i]), colorMap[currentColor]) || len(colorChars) == 0 {
					if strings.HasPrefix(color, "rainbow") {
						colorArray[i] = "\033[3" + strconv.Itoa(i%6+1) + "m"
						break
					}
					colorArray[i] = convertColor(currentColor)
					break
				} else if strings.HasSuffix(currentColor, "+pos") {
					if strings.ContainsAny(strconv.Itoa(i), colorMap[currentColor]) {
						if strings.HasPrefix(color, "rainbow") {
							colorArray[i] = "\033[3" + strconv.Itoa(i%6+1) + "m"
							break
						}
						colorArray[i] = convertColor(currentColor)
						break
					}
				}
			}
		} else {
			colorArray[i] = "\033[0m"
		}
	}
	return colorArray
}

func getColoredCharacters(color string) (map[string]string, string) {
	mappedChars := make(map[string]string)
	allChars := ""
	colorPairs := []string{}
	// Splits the color string into pairs of color:characters
	if strings.Contains(color, ",") {
		colorPairs = strings.Split(color, ",")
	} else {
		// colorPairs = now contains a single pair of color:characters (colorPairs[0])
		colorPairs = append(colorPairs, color)
	}
	// Loop through each set of colors and their character pair
	for _, v := range colorPairs {
		currentColor, characters := "", ""
		// Splits the color:characters pair into color and characters
		if strings.Contains(v, ":") {
			currentColor = strings.Split(v, ":")[0]
			if convertColor(currentColor) == "Invalid color" {
				return nil, ""
			}
			characters = strings.Split(v, ":")[1]
		} else {
			currentColor = v
		}
		// If the characters consist of numbers, then the color is applied to the characters at the specified positions
		if strings.ContainsAny(characters, "0123456789") {
			if strings.Contains(characters, "-") {
				nums := strings.Split(characters, "-")
				num1, num2 := atoi(nums[0]), atoi(nums[1])
				for i := num1; i <= num2; i++ {
					mappedChars[currentColor+"+pos"] += strconv.Itoa(i)
					allChars += strconv.Itoa(i)
				}
			} else {
				println(characters)
				mappedChars[currentColor+"+pos"] += characters
				allChars += characters
			}
		} else {
			// If the characters are not numbers, then the color is applied to the characters
			mappedChars[currentColor] = characters
			allChars += characters
		}
	}
	return mappedChars, allChars
}

func convertColor(color string) string {
	color = strings.TrimSuffix(color, "+pos")
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
	case "rainbow":
		return ""
	default:
		return "Invalid color"
	}
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
