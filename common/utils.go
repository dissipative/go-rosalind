package common

import (
	"bufio"
	"io"
	"log"
	"os"
)

const (
	A = 'A'
	T = 'T'
	G = 'G'
	C = 'C'
)

// ReadInput reads filename into slice of strings
func ReadInput(filename string) (output []string, err error) {
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
