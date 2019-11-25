package main

import "fmt"

func main() {
	reverseString([]byte("hello"))
	reverseString([]byte("hannah"))
}

/*
	数组的reverse操作
*/
func reverseString(s []byte) {
	start := 0
	end := len(s) - 1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
	fmt.Println(string(s))
}
