package main

import (
	"fmt"
	"strconv"
	"strings"
)

func is_palindrome(number int) bool{
	s := strconv.FormatInt(int64(number), 10)
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
	nums := []int{}
	nums_2 := []int{}
	for i := 100; i < 1000; i++ {
		nums = append(nums, i)
		nums_2 = append(nums_2, i)
	}



	answer := 0
	for i := range nums {
		for j := range nums_2 {
			test_value := nums[i] * nums_2[j]
			//fmt.Printf("%d", test_value)
			if is_palindrome(test_value) {
				if test_value > answer {
					answer = test_value
				}

			}
		}
	}
	fmt.Printf("%d", answer)
}
