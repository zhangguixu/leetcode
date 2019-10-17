package main

import "fmt"

func main() {
	s := "leetcode"
	fmt.Println(firstUniqChar(s))
	fmt.Println(firstUniqChar("loveleetcode"))
}

// 这个实现有点慢，看了一下leetcode的solution，执行的速度更慢
func firstUniqChar(s string) int {
	tmpMap := make(map[rune]int, len(s))
	idx := -1

	for _, b := range s {
		tmpMap[b]++
	}

	for i, b := range s {
		if tmpMap[b] == 1 {
			idx = i
			break
		}
	}

	return idx
}
