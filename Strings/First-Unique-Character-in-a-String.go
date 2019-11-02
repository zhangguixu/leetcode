package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))

	fmt.Println(firstUniqChar_v1("leetcode"))
	fmt.Println(firstUniqChar_v1("loveleetcode"))
}

/*
	hashtable two pass
*/
func firstUniqChar(s string) int {
	countMap := make(map[byte]int, len(s))

	for i := 0; i < len(s); i++ {
		countMap[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if countMap[s[i]] == 1 {
			return i
		}
	}

	return -1
}

/*
	array two pass
	题目中假设都是小写的字母，总共就26个，那么可以用一个长度为26的数组，来替代hashtable
*/
func firstUniqChar_v1(s string) int {
	countArr := make([]int, 26)

	for i := 0; i < len(s); i++ {
		countArr[s[i] - 'a']++
	}

	for i := 0; i < len(s); i++ {
		if countArr[s[i] - 'a'] == 1 {
			return i
		}
	}

	return -1
}
