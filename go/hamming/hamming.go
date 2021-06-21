package hamming

import (
	"fmt"
)

func Distance(dnaA, dnaB string) (int, error) {
	if len(dnaA) != len(dnaB) {
		return 0, fmt.Errorf("Strands must be the same length\n%s\n%s", dnaA, dnaB)
	}
	distance := 0
	for i := 0; i < len(dnaA); i++ {
		if dnaA[i] != dnaB[i] {
			distance++
		}
	}
	return distance, nil

}
