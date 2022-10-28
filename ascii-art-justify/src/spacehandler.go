package asciiart

import (
	"fmt"
	"strings"
)

func SpaceHandler(align string, charArray []string, fontArray [95][8]string) []string {
	// Get the terminal size
	_, cols := GetTermSize()
	// Get the length of the string
	length := getLength(charArray)
	// If the string is too long, return nil
	if length > cols {
		fmt.Println("Error: String too long")
		return nil
	}
	// If the string is too short, add spaces to the spaceArray
	if length < cols {
		for i := range charArray {
			switch align {
			case "left":
				for j := 0; j < cols-length; j++ {
					charArray[i] += " "
				}
			case "right":
				for j := 0; j < cols-length; j++ {
					charArray[i] = " " + charArray[i]
				}
				charArray[i] = " " + strings.TrimSuffix(charArray[i], " ")
			case "center":
				for j := 0; j < (cols-length)/2; j++ {
					charArray[i] = " " + charArray[i]
				}
			case "justify":
				// Separate words from charArray and remove spaces, then get new total length without spaces
				words := strings.Fields(strings.Join(charArray, ""))
				length = getLength(words)

			default:
				fmt.Println("Error: Invalid alignment")
				return nil
			}
		}
	}
	return charArray
}

func getLength(charArray []string) int {
	length := 0
	for _, line := range charArray {
		if len(line) > length {
			length = len(line)
		}
	}
	return length
}
