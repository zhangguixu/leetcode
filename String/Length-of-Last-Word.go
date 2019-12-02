package main

import (
	"fmt"
	"regexp"
)

func main() {
	// fmt.Println(lengthOfLastWord("  hello worlsssssd  "))
	fmt.Println(lengthOfLastWord_v1("hello aaaaaaaaaa "))

}

// A word is defined as a character sequence consists of non-space characters only
// 正则表达式
func lengthOfLastWord(s string) int {
	findStr := regexp.MustCompile(`\S+`).FindAllString(s, -1)
	if len(findStr) == 0 {
		return 0
	}
	return len(findStr[len(findStr) - 1])
}

// 比正则表达式更加节省空间
func lengthOfLastWord_v1(s string) int {
	lastWord, lastSpaceIdx, i := "", -1, 0
	for i = 0; i < len(s); i++ {
		if s[i] != ' ' {
			continue
		}
		if lastSpaceIdx + 1 < i {
			lastWord = s[lastSpaceIdx + 1:i]
		}
		lastSpaceIdx = i
	}
	if lastSpaceIdx + 1 < i {
		lastWord = s[lastSpaceIdx + 1:i]
	}
	return len(lastWord)
}
