package asciiart

import (
	"fmt"
)

func AsciiPrint(s string, font string) {
	fontArray := GetFont(font)
	charArray := make([]string, 8)

	// Loop through each character in the string
	for i := 0; i < len(s); i++ {
		// Get the valid character from the font
		if s[i] != '\n' && s[i] >= 32 && s[i] <= 126 {
			for linePos, line := range GetCharacter(rune(s[i]), fontArray) {
				charArray[linePos+len(charArray)-8] += line
			}
		} else if s[i] == '\n' {
			// If it's the first newline in a sequence, add 8 lines, otherwise 1
			charArray = append(charArray, make([]string, 8)...)
			if i < len(s)-1 {
				for j := i; s[j+1] == '\n'; j++ {
					charArray = append(charArray, make([]string, 1)...)
					i++
				}
			}
		} else {
			fmt.Println("ERROR")
		}
	}
	for _, line := range charArray {
		fmt.Println(line)
	}
}
