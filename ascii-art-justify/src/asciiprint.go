// need to figure out length of charArray when fully compiled

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
	// length := assembleLines(s, font)
	// fmt.Println(length, "asciiprint")
	for i := 0; i < len(s); i++ {
		// Get the valid character from the font and store it in the charArray
		if s[i] != '\n' {
			for linePos, line := range GetCharacter(rune(s[i]), fontArray) {
				// if s[i] == ' ' {
				// amountOfWords := strings.Count(s, " ")
				// take in align, charArray, and amountOfWords
				// i++
				// charArray[linePos] += spaceDetector(align, line, amountOfWords, length)
				// charArray[linePos] += spaceDetector(align, charArray[linePos], amountOfWords)
				// }
				charArray[linePos] += line
				// fmt.Println(charArray[linePos])
				// ???
				// charArray[linePos+len(charArray)-AsciiCharHeight] += line
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

func assembleLines(s string, font string) int {
	fontArray, err := GetFont(font)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	charArray := make([]string, 8)
	for i := 0; i < len(s); i++ {
		if s[i] != '\n' && s[i] >= 32 && s[i] <= 126 {
			for linePos, line := range GetCharacter(rune(s[i]), fontArray) {
				charArray[linePos+len(charArray)-8] += line
			}
		} else if s[i] == '\n' {
			if i == 0 || i == len(s)-1 || s[i+1] == '\n' {
				charArray = append(charArray, make([]string, 1)...)
			} else {
				charArray = append(charArray, make([]string, 8)...)
			}
		}
	}
	return len(charArray[0])
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
