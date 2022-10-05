package main

import (
	"fmt"
	"os"
	"strings"

	asciiart "asciiart/src"
)

func main() {
	if len(os.Args[1:]) == 1 {
		asciiart.AsciiPrint(strings.ReplaceAll(os.Args[1], `\n`, "\n"), "standard")
	} else if len(os.Args[1:]) == 2 {
		asciiart.AsciiPrint(strings.ReplaceAll(os.Args[1], `\n`, "\n"), os.Args[2])
	} else {
		fmt.Println("Usage: go run main.go \"[input]\" [font]")
	}
}
