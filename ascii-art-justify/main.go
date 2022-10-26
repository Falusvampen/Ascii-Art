package main

import (
	"fmt"
	"os"
	"strings"

	asciiart "asciiart/src"
)

func main() {
	if len(os.Args[1:]) == 3 && strings.HasPrefix(os.Args[1], "--align=") {
		asciiart.AsciiPrint(os.Args[1][8:], strings.ReplaceAll(os.Args[2], `\n`, "\n"), os.Args[3])
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
	}
}
