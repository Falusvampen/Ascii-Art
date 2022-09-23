package asciiart

import "fmt"

func AsciiPrint(s string, font string) {
	// BUILD CHARACTERS
	fontArray := GetFont(font)
	for _, c := range s {
		for _, line := range GetCharacter(c, fontArray) {
			fmt.Println(line)
		}
	}
}
