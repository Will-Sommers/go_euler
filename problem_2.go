package main

import (
	"fmt"
)

func main() {
	start := 1
	next := 2

	answer := 0
	for next < 4000000 {
		if next % 2 == 0 {
			answer = answer + next
		}
		intermediate := start + next
		start = next
		next = intermediate

	}
	fmt.Printf("%d", answer)
}
