package asciiart

import (
	"fmt"
)

const AsciiCharHeight = 8

func AsciiPrint(align string, s string, font string) []string {
	if s == "dnadiff" {
		DnaDiff()
		return nil
	}
	for _, v := range s {
		if v != '\n' && (v < ' ' || v > '~') {
			fmt.Println("Error: Invalid character")
			return nil
		}
	}
	fontArray, err := GetFont(font)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	charArray := make([]string, 8)
	// ???
	// charArray := initializeLines(s)

	// Loop through each character in the string

	for i := 0; i < len(s); i++ {

		// Get the valid character from the font
		if s[i] != '\n' {
			for linePos, line := range GetCharacter(rune(s[i]), fontArray) {
				charArray[linePos+len(charArray)-AsciiCharHeight] += line
			}
		} else {
			// If there is no character after the newline, add 1 line, otherwise add 8
			if i == 0 || i == len(s)-1 || s[i+1] == '\n' {
				charArray = append(charArray, make([]string, 1)...)
			} else {
				charArray = append(charArray, make([]string, 8)...)
			}
		}
	}
	// Add the spaces to the charArray
	charArray = SpaceHandler(align, charArray)
	for _, line := range charArray {
		// Print final result
		fmt.Println(line)
	}
	return charArray
}

// func initializeLines(s string) []string {
// Initialize the charArray with the correct amount of lines, 0 if a character is invalid, otherwise 8
// charArray := make([]string, 8)
// ????
// for _, c := range s {
// 	if c >= 32 && c <= 126 {
// 		charArray = make([]string, 8)
// 		break
// 	}
// }
// fmt.Println(charArray)
// return charArray
// }
