package asciiart

import "fmt"

func AsciiPrint(s string, font string) {
	// BUILD CHARACTERS
	fontArray := GetFont(font)
	// Print fontArray
	for i := 0; i < len(fontArray); i++ {
		fmt.Println(fontArray[i])
	}
	for _, c := range s {
		for _, line := range GetCharacter(c, fontArray) {
			fmt.Println(line)
		}
	}
}
