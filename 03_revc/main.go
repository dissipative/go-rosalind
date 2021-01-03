package main

import (
	"log"
	"strings"

	"github.com/dissipative/rosalind/common"
)

func main() {
	sequences, err := common.ReadInput("rosalind_revc.txt")
	if err != nil {
		log.Fatal(err)
	}

	for i, seq := range sequences {
		log.Printf("sequence %d: %s", i, revc(seq))
	}
}

func revc(seq string) string {
	var sb strings.Builder

	for _, n := range seq {
		switch n {
		case rune(common.A):
			sb.WriteRune(common.T)
		case rune(common.T):
			sb.WriteRune(common.A)
		case rune(common.G):
			sb.WriteRune(common.C)
		case rune(common.C):
			sb.WriteRune(common.G)
		}
	}

	return reverse(sb.String())
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
