package main

import (
	asciiart "asciiart/src"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Checking if the user has inputted a color option. If they have, it will check if the color option
	// is rainbow. If it is, it will change it to the colors of the rainbow. If not, it will print the
	// string in the color the user has inputted. If the user has not inputted a color option, it will
	// print the usage.

	if len(os.Args[1:]) == 2 && strings.HasPrefix(os.Args[2], "--color=") {
		if os.Args[2] == "--color=rainbow" {
			os.Args[2] = "--color=red,orange,yellow,green,cyan,blue,magenta"
		}
		asciiart.AsciiPrint(strings.ReplaceAll(os.Args[1], `\n`, "\n"), os.Args[2][8:])
	} else {
		if len(os.Args[1:]) == 0 {
			os.Args = append(os.Args, "dnadiff")
		}
		// Array instead of single string for better readability
		usageText := []string{
			"Usage: go run . \"" + os.Args[1] + "\" --color=<input>\n\n",
			"Colors: red, orange, yellow, green, cyan, blue, magenta, white, black, rainbow, 0-255\n\n",
			"Color formats:\n",
			"--color=green:2-5 (positions)\n",
			"--color=blue:abc (characters)\n",
			"--color=yellow:1-2,red:abc (multiple)\n",
		}
		fmt.Println(strings.Join(usageText, ""))
	}
}
