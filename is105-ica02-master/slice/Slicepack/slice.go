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

// Reslice tar en slice og "slicer" den opp på nytt.
// Skal bruke en av "allocate metodene" for å allokere plass i minnet.
func Reslice(slc []byte, lidx int, uidx int) []byte {
	slc = AllocateMake(60)

	return slc
}

// CopySlice ???
func CopySlice(copy1 []byte) []byte {
	copyarr := make([]byte, 60, 110)
	copyarr1 := copy1

	copy(copyarr1, copyarr)

	return copyarr
}
