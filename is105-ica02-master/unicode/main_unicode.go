package main

import (
	"fmt"
	"os"

	unicode "github.com/TeamSolutionsUiA/ica02/is105-ica02-master/unicode/uniPack"
)

func main() {

	arguments := os.Args

	output := unicode.Translate(arguments[1], arguments[2])
	fmt.Println(output)

}
