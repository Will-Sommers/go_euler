package main

import (
	"fmt"
)

func main() {
	sum := 1
	var multiples []int
	for sum < 1000 {
		if (sum % 3 == 0 || sum % 5 == 0) {
			multiples = append(multiples, sum)
		}
		sum += 1
	}

	answer := 0
	for num := range multiples {
		answer = answer + multiples[num]
	}
	 fmt.Printf("%d", answer)
}
