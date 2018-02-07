package main

import (
	"fmt"
	"math/big"
	"strings"
	"strconv"
)

func main() {
	factorial := big.NewInt(1)
	for i:= int64(1); i <= 100; i++ {
		z := big.NewInt(i)
		factorial = z.Mul(big.NewInt(i), factorial);
	}

	digits := strings.Split(factorial.String(), "")
	fmt.Printf("\n", digits)

	answer := 0
	for _, digit := range digits {
		i, _ := strconv.Atoi(digit)
		answer = answer + i
	}
	fmt.Printf("%v", answer)
}