package main

import (
	"log"
	"strings"

	"github.com/dissipative/rosalind/common"
)

func main() {
	sequences, err := common.ReadInput("rosalind_rna.txt")
	if err != nil {
		log.Fatal(err)
	}

	for i, seq := range sequences {
		rna := strings.ReplaceAll(seq, "T", "U")

		log.Printf("sequence %d: %s", i, rna)
	}
}
