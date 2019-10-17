package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(1234))
	fmt.Println(reverse(-6789))
}

func reverse(x int) int {
	max := int(math.Pow(2, 31)) - 1
	min := int(-math.Pow(2, 31))
	hasMinus := false

	if x < 0 {
		hasMinus = true
		x = -x
	}

	num := make([]int, 0, 10)

	for x > 0 {
		num = append(num, x%10)
		x = x / 10
	}

	digits := len(num)
	x = 0

	for i := 0; i < digits; i++ {
		x += num[i] * int(math.Pow10(digits-i-1))
	}

	if hasMinus {
		x = -x
	}

	if x < min || x > max {
		x = 0
	}

	return x
}
