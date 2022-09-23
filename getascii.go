package asciiart

import (
	"bufio"
	"log"
	"os"
)

func GetCharacter(c rune, fontArray []string) []string {
	charStart := int(c) - 32
	var lines []string
	for i := 0; i <= 7; i++ {
		lines = append(lines, fontArray[(charStart)-1:][i])
	}
	return lines
}

func GetFont(font string) []string {
	var fontArray = [94][8]string{}
	file, err := os.Open("../fonts/" + font + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentChar := 0
	for scanner.Scan() {

		if scanner.Text() == "" {
			currentChar++
		}

		fontArray = append(fontArray, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return fontArray

}
