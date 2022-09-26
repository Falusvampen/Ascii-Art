package main

import (
	"log"
	"os"
	"strings"

	"asciiart"
)

func main() {
	if len(os.Args[1:]) == 1 {
		asciiart.AsciiPrint(strings.ReplaceAll(os.Args[1], `\n`, "\n"), "standard")
	} else {
		log.Fatal("Usage: go run main.go \"string\"")
	}
}
