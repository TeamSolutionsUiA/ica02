package unicode

import (
	"fmt"
)

// Kode for Oppgave 4a
func Translate(expression string, language string) string {
	transJp := "\xe2\x80\x9c\xe5\x8c\x97\xe3\x81\xa8\xe5\x8d\x97\xe2\x80\x9d"
	transIs := "\x20\x1C\x6E\x6F\x72\xF0\x75\x72\x20\x6F\x67\x20\x73\x75\xF0\x75\x72\x20\x1D\x0A"

	returnString := ""
	testString := ""
	if expression == "nord og sor" {

		if language == "jp" {
			returnString = transJp
		} else if language == "is" {
			returnString = transIs

		}
		for teller := 0; teller < len(returnString); teller++ {
			testString += fmt.Sprintf("%c", returnString[teller])
		}

	} else {
		testString = "invalid expression! "
	}
	return testString
}

// Kode for Oppgave 4b
func UnicodeCodePointDemo() {
	// Hva er dette?
	// Er det likt på MS Windows og macOS?
	fmt.Println("\xf0\x9F\x98\x80")
	fmt.Println("\xf0\x9F\x98\x97")
	// Demonstrerer at deler av et tegn representert med flere bytes
	// kan ikke tolkes innenfor koden (unicode)
	fmt.Println("\xf0\x9F\x98")
	fmt.Println("\xf0\x9F")
	fmt.Println("\xf0")
}
