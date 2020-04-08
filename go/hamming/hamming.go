package hamming

import "errors"

var errDistance = errors.New("Invalid")

//Distance function is to calculate hamming distance between 2 DNAs
func Distance(a, b string) (int, error) {
	dnaSample1, dnaSample2 := []rune(a), []rune(b)
	if len(dnaSample1) != len(dnaSample2) {
		return 0, errDistance
	}

	var counter = 0
	for indx := range dnaSample1 {
		if dnaSample2[indx] != dnaSample1[indx] {
			counter++
		}
	}
	return counter, nil
}
