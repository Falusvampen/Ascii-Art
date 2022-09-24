package main

import (
	"testing"

	"asciiart"
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

func TestAsciiPrint(t *testing.T) {
	for _, test := range testStrings {
		asciiart.AsciiPrint(test, "standard")
	}
}
