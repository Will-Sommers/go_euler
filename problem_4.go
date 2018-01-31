package main

import (
	"fmt"
	"strconv"
	"strings"
)

func is_palindrome(number int64) bool{
	s := strconv.FormatInt(number, 10)
	chars := strings.Split(s, "")
	string_length := len(chars)
	for i := 0; i < string_length; i++ {
		if (i + 1)> string_length/ 2 {
			return true
		}
		if (chars[i] != chars[string_length - 1 - i]) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("%t", is_palindrome(12131))
	nums := []int{}
	for i := 100; i <= 1000; i++ {
		nums = append(nums, i)
	}
	nums_2 := []int{}
	copy(nums_2, nums)

	for num := range nums {
	}
}
