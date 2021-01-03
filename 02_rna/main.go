package main

import (
	"log"
	"strings"

	"github.com/dissipative/go-rosalind/common"
)

func main() {
	sequences, err := common.ReadInput("rosalind_rna.txt")
	if err != nil {
		log.Fatal(err)
	}

	for i, seq := range sequences {
		log.Printf("sequence %d: %s", i, rna(seq))
	}
}

func rna(seq string) string {
	return strings.ReplaceAll(seq, "T", "U")
}
