package asciiart

import (
	"fmt"
	"os/exec"
	"strconv"
)

func GetTermSize() (lines int, cols int) {
	linesRaw, err := exec.Command("tput", "lines").Output()
	if err != nil {
		return errPrint(err)
	}
	colsRaw, err := exec.Command("tput", "cols").Output()
	if err != nil {
		return errPrint(err)
	}
	lines, err = strconv.Atoi(string(linesRaw))
	if err != nil {
		return errPrint(err)
	}
	cols, err = strconv.Atoi(string(colsRaw))
	if err != nil {
		return errPrint(err)
	}
	return lines, cols
}

func errPrint(err error) (int, int) {
	fmt.Println(err)
	return 0, 0
}