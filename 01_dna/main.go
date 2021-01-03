package main

import (
	"log"
	"strings"

	"github.com/dissipative/rosalind/common"
)

const (
	a = "A"
	t = "T"
	g = "G"
	c = "C"
)

func main() {
	sequences, err := common.ReadInput("rosalind_dna.txt")
	if err != nil {
		log.Fatal(err)
	}

	for i, seq := range sequences {
		countA := strings.Count(seq, a)
		countT := strings.Count(seq, t)
		countG := strings.Count(seq, g)
		countC := strings.Count(seq, c)

		log.Printf("sequence %d: %d %d %d %d", i, countA, countC, countG, countT)
	}

}
