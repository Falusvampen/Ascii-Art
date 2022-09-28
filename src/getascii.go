package asciiart

import (
	"bufio"
	"os"
)

// 95 is the amount of characters in the font and 8 is the amount of lines in the font
func GetFont(font string) ([95][8]string, error) {
	var fontArray = [95][8]string{}
	file, err := os.Open("fonts/" + font + ".txt")
	if err != nil {
		return [95][8]string{}, err
	}
	defer file.Close()

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