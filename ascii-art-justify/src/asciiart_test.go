package asciiart

// var testStrings = []string{
// 	"",
// 	"\n",
// 	"Hello\n",
// 	"hello",
// 	"HeLlO",
// 	"Hello There",
// 	"{Hello There}",
// 	"Hello\nThere",
// 	"Hello\n\nThere",
// }

// // func TestAsciiPrint(t *testing.T) {
// // 	testLines := [9][32]string{}
// // 	file, err := os.Open("test.txt")
// // 	if err != nil {
// // 		fmt.Println(err)
// // 		return
// // 	}

// // 	scanner := bufio.NewScanner(file)
// // 	currentTest := -1
// // 	currentLine := 0
// // 	for scanner.Scan() {
// // 		if strings.ContainsAny(scanner.Text(), "012345678") {
// // 			currentTest++
// // 			currentLine = 0
// // 		} else {
// // 			testLines[currentTest][currentLine] = scanner.Text()
// // 			currentLine++
// // 		}
// // 	}
// // 	if err := scanner.Err(); err != nil {
// // 		file.Close()
// // 		return
// // 	}

// // 	err = file.Close()
// // 	if err != nil {
// // 		fmt.Println(err)
// // 		return
// // 	}
// // 	for i, test := range testStrings {
// // 		println(i)
// // 		lines := AsciiPrint(test, "../fonts/standard.txt")
// // 		for j, line := range lines {
// // 			if line != testLines[i][j] {
// // 				t.Fatalf("Test failed: " + line)
// // 				return
// // 			}
// // 		}
// // 	}
// // }
