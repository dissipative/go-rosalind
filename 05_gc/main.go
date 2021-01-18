package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/dissipative/go-rosalind/common"
	"github.com/shenwei356/bio/seqio/fastx"
)

func main() {
	reader, err := fastx.NewDefaultReader("rosalind_gc.txt")
	if err != nil {
		log.Fatal(err)
	}

	var highestGC float64
	var nameOfHighestGC string
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		gc := countGC(string(record.Seq.Seq))
		if gc > highestGC {
			highestGC = gc
			nameOfHighestGC = string(record.Name)
		}
	}

	fmt.Printf("%s\n%f", nameOfHighestGC, highestGC)
}

func countGC(seq string) float64 {
	seqLen := 100 / float64(len(seq))

	countG := strings.Count(seq, string(common.G))
	countC := strings.Count(seq, string(common.C))

	return seqLen * float64(countG+countC)
}
