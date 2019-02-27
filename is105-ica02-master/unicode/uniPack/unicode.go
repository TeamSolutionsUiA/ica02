package unicode

import (
	"fmt"
)

// Kode for Oppgave 4a
func Translate(expression string, language string) string {
	transJp := "\x20\x1C\x53\x17\x30\x68\x53\x57\x20\x1D\x0A"
	transIs := "\x20\x1C\x6E\x6F\x72\xF0\x75\x72\x20\x6F\x67\x20\x73\x75\xF0\x75\x72\x20\x1D\x0A"

	returnString := ""
	if expression == "nord og sor" {

		if language == "jp" {
			returnString = fmt.Sprintf("%x", transJp)
		} else if language == "is" {
			returnString = fmt.Sprintf("%x", transIs)

		}

	} else {
		returnString = "invalid expression! "
	}
	return returnString
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
