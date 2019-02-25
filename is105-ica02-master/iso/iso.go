package iso

import "fmt"

// Oppgave 2a
// Lag selv en "string literal" med utvidet ASCII
// Blir det kun en slik "string literal" eller trenger man flere
// for å representere utvidet ASCII?

func GetExtendedASCIIStringLiteral() string {
	stringLitteral := ""
	for i := 128; i < 255; i++ {
		stringLitteral += fmt.Sprintf("%X", i)
	}
	return stringLitteral
}

// IterateOverASCIIStringLiteral tar en "string literal" som INN-data
func IterateOverExtendedASCIIStringLiteral(stringLitteral string) {
	// Kode for Oppgave 2a
	for i := 0; i < len(stringLitteral); i++ {
		fmt.Printf("%X  ", stringLitteral[i])
		fmt.Printf("%q  ", stringLitteral[i])
		fmt.Printf("%b  \n", stringLitteral[i])
	}
}

// GreetingExtendedASCII returnerer en tekst-streng i
// utvidet ASCII
// Kode for Oppgave 2c
func GreetingExtendedASCII() string {
	return ""
}
