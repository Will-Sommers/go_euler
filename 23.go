package main

import (
	"fmt"
	"runtime"
	"sort"
)

const MAX = 30000

func divmod(num, denom int) (quotient, remainder int) {
	quotient = num / denom
	remainder = num % denom
	return
}

func divisors(i int) []int {
	nums := []int{}
	for i:=2; i < MAX; i++ {
		nums = append(nums, i)
	}
	data := []int{1}
	for _, num := range nums {
		_, remainder := divmod(i, num)
		if remainder == 0 && num != i {
			data = append(data, []int{num}...)
		}
		if (num >= (i / 2 )) {
			break
		}
	}
	return data
}

func abundant(i int, divisors []int) bool {
	abundant := false;
	sum_of_divisors := 0
	for _, v := range divisors {
		sum_of_divisors += v
		if sum_of_divisors > i {
			abundant = true
			break
		}
	}
	return abundant
}


func main() {
	runtime.GOMAXPROCS(2)
	var slice []int
	queue := make(chan int, 1)
	for i := 1; i <= MAX; i++ {
		go func(i int) {
			is_abundant := abundant(i, divisors(i))
			if is_abundant {
				queue <- i
			} else {
				queue <- -1
			}
		}(i)
	}

	remaining := MAX
	for t := range queue {
		slice = append(slice, t)
		if remaining--; remaining == 0 {
			close(queue)
		}
	}

	var clean_slice []int
	for _, v := range slice {
		if v > 0 {
			clean_slice = append(clean_slice, v)
		}
	}
	sort.Ints(clean_slice)
	fmt.Println(clean_slice)
	var answer []int
	for i := 1; i <= MAX; i++ {
		answer = append(answer, i)
	}
	for j := range clean_slice {
		for i := range clean_slice {
			sum := clean_slice[i] + clean_slice[j]
			if sum == 24 {
				fmt.Println(clean_slice[i], clean_slice[j])
			}
			if sum <= (len(answer)) {
				answer[sum - 1] = -1
			}
		}
	}

	answer_int := 0

	fmt.Println(answer)
	for _, v := range answer {
		if v > 0 {
			answer_int += v
		}
	}

	fmt.Println(answer_int)
}