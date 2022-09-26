package asciiart

import (
	"fmt"
)

func AsciiPrint(s string, font string) {
	if s == "\n" {
		fmt.Println()
		return
	} else if s == "dnadiff" {
		DnaDiff()
		return
	}
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
			// If there is no character after the newline, add 1 line, otherwise add 8
			if i < len(s)-1 && s[i+1] == ' ' || i == len(s)-1 || s[i+1] == '\n' {
				charArray = append(charArray, make([]string, 1)...)
				println("one line")
			} else {
				charArray = append(charArray, make([]string, 8)...)
				println("eight lines")
			}
			// If newlines are found in a sequence after the initial, add one line for each
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
