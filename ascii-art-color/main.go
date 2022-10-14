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
		fmt.Println("Usage: go run . \"" + os.Args[1] + "\" --color=<input>\n\nColor formats:\n--color=green:2-5 (positions)\n--color=blue:abc (letters)\n--color=yellow:1-2,red:abc (multiple)")
	}
}
