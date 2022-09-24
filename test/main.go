package main

import (
	"log"
	"os"
	"strings"

	"asciiart"
)

func main() {
	if len(os.Args[1:]) == 1 {
		asciiart.AsciiPrint(strings.Replace(os.Args[1], `\n`, "\n", -1), "standard")
	} else {
		log.Fatal("Usage: go run main.go \"string\"")
	}
}
