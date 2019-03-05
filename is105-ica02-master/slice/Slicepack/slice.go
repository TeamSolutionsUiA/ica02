package slice

//import "fmt"

// AllocateVar har INN-argument b
// b - antall bytes brukeren ønsker å allokere 
// Returnerer en slice av type []byte
// 
func AllocateVar(b int) []byte {
	var newSlice []byte
	newSlice = make([]byte, b, b)
	return newSlice
}

// AllocateMake tar lengde og kapasitet som b og lager en ny slice
//
func AllocateMake(b int) []byte {
	newSlice := make([]byte, b, b)
	return newSlice
}

// Reslice takes a slice and reslices it
func Reslice(slc []byte, lidx int, uidx int) []byte {
	slc =make([]byte, lidx, uidx)

	return slc
}

// CopySlice ???
func CopySlice (copy1 []byte) []byte {
copyarr :=make([]byte, 5,5)
copy1 =make([]byte,20,20)	
copy(copy1, copyarr)

return copyarr
}