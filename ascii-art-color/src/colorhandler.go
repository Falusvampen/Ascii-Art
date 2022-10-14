package asciiart

import (
	"strconv"
	"strings"
)

// It splits the color string into pairs of color:characters, sorts the pairs, and then loops through
// each pair to set the color of the characters
func ColorHandler(s string, color string) []string {
	// Initialize array to store color positions
	colorArray := make([]string, len(s))

	// Splits the color string into pairs of color:characters
	colorPairs := strings.Split(color, ",")
	colorPairs = sortColorPairs(colorPairs)

	// Loop through each set of colors and their character pair
	for _, pair := range colorPairs {
		currentColor, characters := "", ""

		// Splits the color:characters pair into color and characters
		if strings.Contains(pair, ":") {
			currentColor = strings.Split(pair, ":")[0]
			if convertColor(currentColor) == "Invalid color" {
				return nil
			}
			characters = strings.Split(pair, ":")[1]
		} else {
			currentColor = pair
			if convertColor(currentColor) == "Invalid color" {
				return nil
			}
		}

		if characters == "" {
			// If no characters are specified, color the entire string
			colorArray = setFullColor(s, currentColor, colorArray, colorPairs)
		} else {
			// If characters are specified, color them
			colorArray = setSpecificColor(s, characters, currentColor, colorArray)
		}
	}

	return finalizeColor(colorArray)
}

// It sorts the color pairs from full colors to specific colors
func sortColorPairs(colorPairs []string) []string {
	// Sort the color pairs from full colors to specific colors
	fullColors := []string{}
	specificColors := []string{}
	for _, pair := range colorPairs {
		if strings.Contains(pair, ":") {
			specificColors = append(specificColors, pair)
		} else {
			fullColors = append(fullColors, pair)
		}
	}
	return append(fullColors, specificColors...)
}

// Add the reset color to every empty position
func finalizeColor(colorArray []string) []string {
	// Add the reset color to every empty position
	for i := range colorArray {
		if colorArray[i] == "" {
			colorArray[i] = "\033[0m"
		}
	}
	return colorArray
}

func setFullColor(s string, currentColor string, colorArray []string, colorPairs []string) []string {
	// Get all colorPairs that do not include ':' and put them into a new colorPairs
	fullColors := []string{}
	for _, pair := range colorPairs {
		if !strings.Contains(pair, ":") {
			fullColors = append(fullColors, pair)
		}
	}
	// If the characters are empty, then the color is applied to the entire string
	for i := range s {
		if len(fullColors) > 1 {
			// If there are multiple colors, then the color is based on the position of the character
			colorArray[i] = convertColor(fullColors[i%len(fullColors)])
		} else {
			colorArray[i] = convertColor(currentColor)
		}
	}
	return colorArray
}

// If the characters are a range of positions, then the color is applied to the characters at the
// specified positions, otherwise the color is applied to the specific characters
func setSpecificColor(s string, characters string, currentColor string, colorArray []string) []string {
	// If the characters consist of numbers, then the color is applied to the characters at the specified positions
	if strings.ContainsAny(characters, "0123456789") && strings.Contains(characters, "-") && !strings.ContainsAny(characters, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		nums := strings.Split(characters, "-")
		num1, num2 := atoi(nums[0])-1, atoi(nums[1])-1
		// Prevent num2 from going out of bounds
		if num2 > len(s) {
			num2 = len(s) - 1
		}
		for i := num1; i <= num2; i++ {
			colorArray[i] = convertColor(currentColor)
		}
	} else {
		// If the characters are not a range of positions, then the color is applied to the specific characters
		for i := range s {
			if strings.Contains(characters, string(s[i])) {
				colorArray[i] = convertColor(currentColor)
			}
		}
	}
	return colorArray
}

// If the color is a number, return the color code for that number. If the color is a string, return
// the color code for that string
func convertColor(color string) string {
	if strings.ContainsAny(color, "0123456789") && !strings.Contains(color, "-)") && !strings.ContainsAny(color, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return "\033[38;5;" + color + "m"
	}
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
	case "orange":
		return "\033[38;5;208m"
	default:
		return "Invalid color"
	}
}

// It takes a string and returns an integer

func atoi(s string) int { // our little pet
	i, _ := strconv.Atoi(s)
	return i
}
