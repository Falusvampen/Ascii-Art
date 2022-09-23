package asciiart

func AsciiPrint(s string, font string){
	// BUILD CHARACTERS
	for _, c := range s {
		GetCharacter(c, font)
	}
	
}
