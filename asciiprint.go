package asciiart

import "fmt"

func AsciiPrint(s string, font string) {
	// BUILD CHARACTERS
	fontArray := GetFont(font)
	charArray := []string{}
	for _, c := range s {
		for i, line := range GetCharacter(c, fontArray) {
			charArray[i] += line
		}
	}
	for _, line := range charArray {
		fmt.Println(line)
	}
}

func initalizeLines(s string) []string {

	//charArray := make([]string, len(s))
}
