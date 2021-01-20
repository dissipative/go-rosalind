package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/dissipative/go-rosalind/common"
)

func main() {
	sequences, err := common.ReadInput("rosalind_hamm.txt")
	if err != nil {
		log.Fatal(err)
	}
	if len(sequences) != 2 {
		log.Fatal("there must be 2 sequences on input")
	}

	diff, err := hamm(sequences[0], sequences[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(diff)
}

func hamm(seq1, seq2 string) (int, error) {
	diff := 0

	if len(seq1) != len(seq2) {
		return diff, errors.New("sequences must be equal length")
	}

	for i, n := range seq1 {
		if n != rune(seq2[i]) {
			diff++
		}
	}

	return diff, nil
}
