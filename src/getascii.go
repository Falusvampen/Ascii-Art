package asciiart

import (
	"bufio"
	"os"
)

// 95 is the amount of characters in the font and 8 is the amount of lines in the font
func GetFont(font string) ([95][8]string, error) {
	var fontArray = [95][8]string{}

	// If font argument is a shortcut, provide it with a path
	fonts := []string{"standard", "shadow", "thinkertoy"}
	for _, f := range fonts {
		if font == f {
			font = "fonts/" + f + ".txt"
			break
		}
	}

	file, err := os.Open(font)
	if err != nil {
		return [95][8]string{}, err
	}

	scanner := bufio.NewScanner(file)
	currentChar := -1
	currentLine := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			currentChar++
			currentLine = 0
		} else {
			fontArray[currentChar][currentLine] = scanner.Text()
			currentLine++
		}
	}
	if err := scanner.Err(); err != nil {
		file.Close()
		return [95][8]string{}, err
	}

	err = file.Close()
	return fontArray, err
}

func GetCharacter(c rune, fontArray [95][8]string) []string {
	charStart := int(c) - 32
	var lines []string
	for i := 0; i <= 7; i++ {
		lines = append(lines, fontArray[charStart][i])
	}
	return lines
}
