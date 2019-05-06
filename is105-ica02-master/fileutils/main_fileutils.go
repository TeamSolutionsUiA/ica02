package main

import (
	"fmt"

	"github.com/TeamSolutionsUiA/ica02/is105-ica02-master/fileutils/fileutils"
)

func main() {
	lang01 := fileutils.FileToByteslice("../files/lang01.wl")
	lang02 := fileutils.FileToByteslice("../files/lang02.wl")
	lang03 := fileutils.FileToByteslice("../files/lang03.wl")
	fmt.Printf("%c \n%c \n%c", lang01, lang02, lang03)
}
