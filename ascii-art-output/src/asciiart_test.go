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

func TestAsciiPrint(*testing.T) {
	for _, v := range testStrings {
		AsciiPrint(v, "../fonts/standard.txt")
	}
}
