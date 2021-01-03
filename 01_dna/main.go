package main

import (
	"log"
	"strings"

	"github.com/dissipative/go-rosalind/common"
)

func main() {
	sequences, err := common.ReadInput("rosalind_dna.txt")
	if err != nil {
		log.Fatal(err)
	}

	for i, seq := range sequences {
		countA, countC, countG, countT := countDNANucleotides(seq)
		log.Printf("sequence %d: %d %d %d %d", i, countA, countC, countG, countT)
	}
}

func countDNANucleotides(seq string) (countA, countC, countG, countT int) {
	countA = strings.Count(seq, string(common.A))
	countT = strings.Count(seq, string(common.T))
	countG = strings.Count(seq, string(common.G))
	countC = strings.Count(seq, string(common.C))

	return
}
