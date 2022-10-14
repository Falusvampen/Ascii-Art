package asciiart

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

// Go test for testing color
var testColors = []string{
	"red",
	"orange",
	"yellow",
	"green",
	"cyan",
	"blue",
	"magenta",
	"white",
	"black",
	"rainbow",
}

// TestColor tests the color function
func TestColor(t *testing.T) {
	// Go through the testColors array and test each color
	for _, color := range testColors {
		// If the color is rainbow, test it by replacing it with the colors of the rainbow
		if color == "rainbow" {
			color = "red,orange,yellow,green,cyan,blue,magenta"
		}

		// Run the color function and compare the output to the expected output
		output := exec.Command("go", "run", ".", "Hello, world!", "--color="+color)
		output.Dir = ".."
		outputString, _ := output.Output()

		fmt.Println(string(outputString))
		for _, currentColor := range strings.Split(color, ",") {
			if !strings.Contains(string(outputString), convertColor(currentColor)) {
				t.Error("033[31mColor test failed: " + currentColor)
			} else {
				fmt.Println("\033[32mColor test passed: " + convertColor(currentColor) + currentColor + "\033[0m")
			}
		}
	}
}
