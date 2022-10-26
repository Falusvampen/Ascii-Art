package asciiart

import "fmt"

func SpaceHandler(align string, s string, fontArray [95][8]string) []string {
	// Get the terminal size
	_, cols := GetTermSize()
	// Get the length of the string
	length := 0
	for _, c := range s {
		maxLength := 0
		for _, line := range GetCharacter(c, fontArray) {
			if len(line) > maxLength {
				maxLength = len(line)
			}
		}
		length += maxLength
	}
	// Initialize the spaceArray with the correct amount of spaces
	spaceArray := make([]string, length)
	// If the string is too long, return nil
	if length > cols {
		fmt.Println("Error: String too long")
		return nil
	}
	// If the string is too short, add spaces to the spaceArray
	if length < cols {
		switch align {
		case "left":
			for i := 0; i < cols-length; i++ {
				spaceArray = append(spaceArray, " ")
			}
		case "right":
			for i := 0; i < cols-length; i++ {
				spaceArray = append([]string{" "}, spaceArray...)
			}
		case "center":
			for i := 0; i < (cols-length)/2; i++ {
				spaceArray = append([]string{" "}, spaceArray...)
			}
			for i := 0; i < (cols-length)/2; i++ {
				spaceArray = append(spaceArray, " ")
			}
		default:
			return nil
		}
	}
	return spaceArray
}
