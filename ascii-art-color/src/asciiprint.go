package asciiart

import (
	"fmt"
)

// It takes a string and a color string, and returns an array of strings that are the ascii art of the
// string
func AsciiPrint(s string, colorData string) []string {
	if s == "dnadiff" {
		DnaDiff()
		return nil
	}
	fontArray, err := GetFont("standard")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	charArray := initializeLines(s)
	colorArray := ColorHandler(s, colorData)
	if colorArray == nil {
		fmt.Println("Error: Invalid color input")
		return nil
	}

	// Loop through each character in the string
	for i := 0; i < len(s); i++ {
		// Get the valid character from the font
		if s[i] != '\n' && s[i] >= 32 && s[i] <= 126 {
			for linePos, line := range GetCharacter(rune(s[i]), fontArray) {
				charArray[linePos+len(charArray)-8] =
					charArray[linePos+len(charArray)-8] + colorArray[i] + line
			}
		} else if s[i] == '\n' {
			// If there is no character after the newline, add 1 line, otherwise add 8
			if i == 0 || i == len(s)-1 || s[i+1] == '\n' {
				charArray = append(charArray, make([]string, 1)...)
			} else {
				charArray = append(charArray, make([]string, 8)...)
			}
		} else {
			fmt.Println("Error: Invalid character")
			return nil
		}
	}
	for _, line := range charArray {
		fmt.Println(line)
	}
	return charArray
}

func initializeLines(s string) []string {
	// Initialize the charArray with the correct amount of lines, 0 if a character is invalid, otherwise 8
	charArray := make([]string, 0)
	for _, c := range s {
		if c >= 32 && c <= 126 {
			charArray = make([]string, 8)
			break
		}
	}
	return charArray
}
