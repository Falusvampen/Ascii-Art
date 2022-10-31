package asciiart

import (
	"fmt"
	"strings"
)

func SpaceHandler(align string, charArray []string) []string {

	// Get the terminal size
	_, cols := GetTermSize()

	// Get the length of the string
	// ??
	// length := getLength(charArray)
	length := len(charArray[0])
	// fmt.Println(len(charArray[0]))

	// getWords(charArray)
	// If the string is too long, return nil
	if length > cols {
		fmt.Println("Error: String too long")
		return nil
	}
	// If the string is too short, add spaces to the spaceArray
	if length < cols {
		for i := range charArray {
			switch align {
			// Left is already aligned
			case "left":
			// ???
			// for j := 0; j < cols-length; j++ {
			// 	charArray[i] += " "
			// }
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
				charArray[i] = strings.ReplaceAll(charArray[i], "      ", "#")
			// charArray[i] = spaceDetector(charArray[i])
			default:
				fmt.Println("Error: Invalid alignment")
				return nil
			}
		}
	}
	return charArray
}

// func spaceDetector(line string) string {
// 	_, cols := GetTermSize()
// 	line = strings.Replace(line, "       ", "#", 1)
// 	strLen := len(line)
// 	nice := cols - strLen
// 	spaces := ""
// 	for i := 0; i < nice; i++ {
// 		spaces += " "
// 	}
// 	// line = strings.Replace(line, "#", spaces, 1)
// 	return line
// }
// take in two strings and one int as arguments

func spaceDetector(align string, line string, amountOfWords int, length int) string {

	if align == "justify" {
		_, lines := GetTermSize()
		nice := (lines - length) / amountOfWords
		spaces := ""
		for i := 0; i < nice; i++ {
			spaces += " "
		}
		return line
	} else {
		return line
	}
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

// Not needed
// func getLength(charArray []string) int {
// 	length := 0
// 	for _, line := range charArray {
// 		if len(line) > length {
// 			length = len(line)
// 		}
// 	}
// 	return length
// }
