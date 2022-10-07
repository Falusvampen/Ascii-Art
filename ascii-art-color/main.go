package main

import (
	asciiart "asciiart/src"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args[1:]) == 2 && strings.HasPrefix(os.Args[2], "--color=") {
		asciiart.AsciiPrint(strings.ReplaceAll(os.Args[1], `\n`, "\n"), os.Args[2][8:])
	} else {
		fmt.Println("Usage: go run . [STRING] [OPTION]\n\nEX: go run . something --color=<color>")
	}
}
