package main

import ("github.com/TeamSolutionsUiA/ica02/is105-ica02-master/slice/slicepack"
 "fmt"
)
func main() {
	byteslice1 := make([]byte, 50, 100)

	fmt.Println("&byteslice1[0]")
	fmt.Println(&byteslice1[0])
	aslice := slice.Reslice(byteslice1, 50, 100)
	fmt.Println("&aslice[0]")
	fmt.Println(&aslice[0])
	fmt.Println("&byteslice1[50]")
	fmt.Println(&byteslice1[50])
	

}
