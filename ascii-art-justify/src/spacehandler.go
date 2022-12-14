package asciiart

import (
	"fmt"
	"strings"
)

func SpaceHandler(align string, charArray []string) []string {
	// Get the terminal size
	_, cols := GetTermSize()
	// Get the length of the string
	length := getLength(charArray)
	getWords(charArray)
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

			default:
				fmt.Println("Error: Invalid alignment")
				return nil
			}
		}
	}
	return charArray
}

func getWords(charArray []string) [][]string {
	words := make([][]string, len(charArray))

	// Loop through each column in the charArray
	lastWordEnd := 0
	for row := range charArray {
		lastWordEnd = strings.Index(charArray[5], "       ")
		if lastWordEnd == -1 {
			words[row] = append(words[row], charArray[row])
		} else if len(words) < 9 {
			words[row] = append(words[row], charArray[row][:lastWordEnd])
		} else {
			words[row] = append(words[row], charArray[row][lastWordEnd+7:])
		}
	}

	// Print each word in words.
	for _, word := range words {
		fmt.Println(word)
	}
	return words
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
