package main

import "fmt"

// 位运算处理： https://lihaoquan.me/2018/1/1/bit-operator.html
// https://zh.wikipedia.org/wiki/%E4%BD%8D%E6%93%8D%E4%BD%9C
func main() {
	a := 0
	a ^= 2
	fmt.Println(a)
	a ^= 4
	fmt.Println(a)
	a ^= 2
	fmt.Println(a)
}

func singleNumber(nums []int) int {
	r := 0
	for _, n := range nums {
		r ^= n
	}
	return r
}
