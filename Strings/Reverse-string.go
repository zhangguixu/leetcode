package main

import "fmt"

func main() {
	reverseString([]byte("hello"))
	reverseString([]byte("hannah"))
}

// 不能申请额外的空间，下标交换即可
// 貌似这样执行的速度并不快？
func reverseString(s []byte)  {
	start := 0
	end := len(s) - 1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}

	fmt.Println(string(s))
}