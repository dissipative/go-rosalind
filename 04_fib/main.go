package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/dissipative/go-rosalind/common"
)

func main() {
	samples, err := common.ReadInput("rosalind_fib.txt")
	if err != nil {
		log.Fatal(err)
	}

	for i, sample := range samples {
		input := strings.Fields(sample)
		n, k, err := inputToInt(input)
		if err != nil {
			log.Fatalf("sample %d: %v", i, err)
		}

		log.Printf("sample %d: %d", i, reccurent(n, k))
	}

}

func inputToInt(input []string) (n, k int, err error) {
	if len(input) != 2 {
		err = fmt.Errorf("invalid input: %v", input)
		return
	}

	n, err = strconv.Atoi(input[0])
	if err != nil || n > 40 {
		err = fmt.Errorf("invalid n: [%s]", input[0])
		return
	}
	k, err = strconv.Atoi(input[1])
	if err != nil || k > 5 {
		err = fmt.Errorf("invalid k: [%s]", input[1])
		return
	}

	return
}

func reccurent(n, k int) int {
	f1 := 1
	f2 := 1
	fn := f2

	for i := 3; i <= n; i++ {
		f2 = fn
		fn = f1*k + f2
		f1 = f2
	}

	return fn
}
