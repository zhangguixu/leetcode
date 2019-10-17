package main

import "fmt"

func main() {
	fmt.Println(reverseString([]byte("hello")))
	fmt.Println(reverseString([]byte("abcd")))
}

// 要求不能额外申请空间
func reverseString(s []byte) string {
	i := 0
	l := len(s) - 1
	for i <= l {
		s[i], s[l] = s[l], s[i]
		i++
		l--
	}
	return string(s[:])
}
