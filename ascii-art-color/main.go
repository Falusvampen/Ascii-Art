package main

import (
	asciiart "asciiart/src"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args[1:]) == 2 && strings.HasPrefix(os.Args[2], "--color=") {
		if os.Args[2] == "--color=rainbow" {
			os.Args[2] = "--color=red,orange,yellow,green,cyan,blue,magenta"
		}
		asciiart.AsciiPrint(strings.ReplaceAll(os.Args[1], `\n`, "\n"), os.Args[2][8:])
	} else {
		if len(os.Args[1:]) == 0 {
			os.Args = append(os.Args, "dnadiff")
		}
		fmt.Println("Usage: go run . \"" + os.Args[1] + "\" --color=<input>\n\nColor formats:\n--color=green:2-5 (positions)\n--color=blue:abc (letters)\n--color=yellow:1-2,red:abc (multiple)")
	}
}
