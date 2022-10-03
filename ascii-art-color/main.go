package main

import (
	"fmt"
	"os"
	"strings"

	asciiart "asciiart/ascii-art-color/src"
)

func main() {
	SetColor("red")
	if len(os.Args[1:]) == 2 && strings.HasPrefix(os.Args[3], "--color=") {
		asciiart.AsciiPrint(strings.ReplaceAll(os.Args[1], `\n`, "\n"), os.Args[2])
	} else {
		fmt.Println(
			"Usage: go run . [STRING] [OPTION]\n\nEX: go run . something --color=<color>")
	}
}

func SetColor(color string) {
	switch color {
	case "red":
		fmt.Print("\033[31m")
	case "green":
		fmt.Print("\033[32m")
	case "yellow":
		fmt.Print("\033[33m")
	case "blue":
		fmt.Print("\033[34m")
	case "magenta":
		fmt.Print("\033[35m")
	case "cyan":
		fmt.Print("\033[36m")
	case "white":
		fmt.Print("\033[37m")
	}
}
