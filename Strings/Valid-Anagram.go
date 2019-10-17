package main

import "fmt"

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "cat"))
	fmt.Println(isAnagram("ab", "b"))
	s := "1234"
	fmt.Println(s[1])
}

// 这种做法的运行效率太低了
// func isAnagram(s string, t string) bool {
// 		if len(s) != len(t) {
// 			return false
// 		}
// 		countM := make(map[rune]int, len(s))
// 		flag := true
// 		for _, b := range s {
// 			countM[b]++
// 		}
// 		for _, b := range t {
// 			countM[b]--
// 		}
// 		for _, v := range countM {
// 			if v != 0 {
// 				flag = false
// 				break
// 			}
// 		}
// 		return flag
// }

// 运行效率目前最高的解法
func isAnagram(s string, t string) bool {
	if len(t) != len(s) {
			return false
	}
	
	counter := [26]int{}
	for i := 0; i < len(s); i++ {
			counter[s[i] - 'a']++
			counter[t[i] - 'a']--
	}
	
	for _, count := range counter {
			if count != 0 {
					return false
			}
	}
	
	return true
}