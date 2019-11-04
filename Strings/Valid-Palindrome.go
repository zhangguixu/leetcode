package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(".,"))

	fmt.Println("==================")

	fmt.Println(isPalindrome_v1("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome_v1("race a car"))
	fmt.Println(isPalindrome_v1(".,"))
}

/*
	根据题目的要求：
	- considering only alphanumeric characters and ignoring cases
	- define empty string as valid palindrome
*/

func isAlphanumeric(b byte) bool {
	if (b >= '0' && b <= '9') || (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
		return true
	}
	return false
}

func isEqual(b1 byte, b2 byte) bool {
	if b1 >= 'A' && b2 >= 'A' {
		return b1 == b2 || b1 - 32 == b2 || b1 + 32 == b2
	}
	return b1 == b2
}

/*
	two pointer, one pass 
*/
func isPalindrome(s string) bool {
	start := 0
	tail := len(s) - 1
	for start < tail {
		if !isAlphanumeric(s[start]) {
			start++
			continue
		}
		if !isAlphanumeric(s[tail]) {
			tail--
			continue
		}
		if !isEqual(s[start], s[tail]) {
			return false
		}
		start++
		tail--
	}
	return true
}

/*
	https://leetcode.com/explore/featured/card/top-interview-questions-easy/127/strings/883/discuss/39981/My-three-line-java-solution
	在discuss中有一个思路是

	将字符串中，除了Alphanumeric的字符之外，其他都替换成空格，然后将其与自身的反转作对比

	这个做法无疑在性能上并不是一个最优的选择
	1）正则匹配的运算开销
	2）一个是字符的反转之后，产生的新的存储开销

	但是这个方法提供了一个思路：

		字符串的回文确实是一个字符串 == 一个字符串的反转

	下面就用go语言内置的来实现
*/
func reverse(s string) string {
	n := len(s)
	sList := []byte(s)
	for i := 0; i < n / 2; i++ {
		sList[i], sList[n - i - 1] = sList[n - i - 1], sList[i]
	}
	return string(sList)
}

func isPalindrome_v1(s string) bool {
	reg := regexp.MustCompile(`[^0-9A-Za-z]`)
	str := reg.ReplaceAllString(s, "")
	str = strings.ToLower(str)
	return reverse(str) == str
}
