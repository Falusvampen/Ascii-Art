package asciiart

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetTermSize() (lines int, cols int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	sizeRaw, err := cmd.Output()
	if err != nil {
		return errPrint(err)
	}
	size := strings.Split(string(sizeRaw), " ")
	lines, err = strconv.Atoi(string(strings.TrimSuffix(string(size[0]), "\n")))
	if err != nil {
		return errPrint(err)
	}
	cols, err = strconv.Atoi(string(strings.TrimSuffix(string(size[1]), "\n")))
	if err != nil {
		return errPrint(err)
	}
	return lines, cols
}

func errPrint(err error) (int, int) {
	fmt.Println(err)
	return 0, 0
}
