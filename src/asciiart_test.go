package asciiart

import (
	"testing"
)

var testStrings = []string{
	"",
	"\n",
	"Hello\n",
	"hello",
	"HeLlO",
	"Hello There",
	"{Hello There}",
	"Hello\nThere",
	"Hello\n\nThere",
}

var testAnswers = [9][]string{
	{""},
	{
		" _    _          _   _         ",
		"| |  | |        | | | |         ",
		"| |__| |   ___  | | | |   ___   ",
		"|  __  |  / _ \\ | | | |  / _ \\  ",
		"| |  | | |  __/ | | | | | (_) | ",
		"|_|  |_|  \\___| |_| |_|  \\___/  ",
	},
}

func TestAsciiPrint(*testing.T) {
	for _, test := range testStrings {
		AsciiPrint(test, "standard")
	}

	for _, v := range testAnswers {
		for _, line := range v {
			println(line)
		}
	}
}
