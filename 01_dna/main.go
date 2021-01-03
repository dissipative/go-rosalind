package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

const (
	a = "A"
	t = "T"
	g = "G"
	c = "C"
)

func main() {
	sequences, err := readInput("rosalind_dna.txt")
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

func readInput(filename string) (output []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if line != "" {
			output = append(output, line)
		}

		if err != nil {
			break
		}
	}
	if err != io.EOF {
		return
	}

	return output, nil
}
