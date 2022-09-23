package asciiart

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func GetCharacter(c rune, font string) {
	// Get from font file
	file, err := os.Open("../fonts/" + font + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getCharacters