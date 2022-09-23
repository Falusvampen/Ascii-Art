package asciiart

import (
	"bufio"
	"log"
	"os"
)

func GetCharacter(c rune, fontArray []string) []string {
	var lines []string
	for i := 0; i <= 7; i++ {
		lines = append(lines, fontArray[299-1:][i])
	}
	return lines
}

func GetFont(font string) []string {
	var fontArray []string
	file, err := os.Open("../fonts/" + font + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fontArray = append(fontArray, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return fontArray
}
