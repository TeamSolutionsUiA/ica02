package main

import (
	"github.com/TeamSolutions/ica02/is105-ica02-master/iso"
)

func main() {
	stringliteral := iso.GetExtendedASCIIStringLiteral()
	iso.IterateOverExtendedASCIIStringLiteral(stringliteral)
	iso.GreetingExtendedASCII()

}
