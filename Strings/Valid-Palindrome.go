package main

import "fmt"

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(".,"))
}

/*
	根据题目的要求：
	- considering only alphanumeric characters and ignoring cases
	- define empty string as valid palindrome
	在编写代码的时候要注意
*/

// 这种实现要两次循环，而且有多余的存储，运行效率只击败70%的用户的
// func isPalindrome(s string) bool {
// 	if len(s) == 0 {
// 		return true
// 	}
// 	list := make([]byte, 0, len(s))
// 	for i := 0; i < len(s); i++ {
// 		if (s[i] >= 48 && s[i] <= 57) || (s[i] >= 65 && s[i] <= 90) || (s[i] >= 97 && s[i] <= 122) {
// 			if s[i] >= 97 {
// 				list = append(list, s[i] - 32)
// 			} else {
// 				list = append(list, s[i])
// 			}
// 		}
// 	}
// 	if len(list) == 0 {
// 		return true
// 	}
// 	m := len(list)/2
// 	start := 0
// 	end := len(list) - 1
// 	flag := true
// 	for start < m {
// 		if list[start] != list[end] {
// 			flag = false
// 			break
// 		}
// 		start++
// 		end--
// 	}
// 	return flag
// }

// 只用一次循环，效率目前排行是最高的
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	start := 0
	tail := len(s) - 1
	flag := true
	for start < tail {
		for start <= tail && !isAlphanumeric(s[start]) {
			start++
		}
		for start <= tail && !isAlphanumeric(s[tail]) {
			tail--
		}
		if start >= tail {
			break
		}
		if s[start] == s[tail] || (s[start] >= 'A' && s[tail] >= 'A' && (s[start]-s[tail] == 32 || s[tail]-s[start] == 32)) {
			start++
			tail--
		} else {
			flag = false
			break
		}
	}
	return flag
}

func isAlphanumeric(b byte) bool {
	for (b >= '0' && b <= '9') || (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
		return true
	}
	return false
}
