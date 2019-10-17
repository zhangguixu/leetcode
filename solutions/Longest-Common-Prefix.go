package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"testsss", "test", "testaas"}))
}

func longestCommonPrefix(strs []string) string {

	// 边界处理
	if len(strs) == 0 {
		return ""
	}

	var shortestStr string
	var commonPrefix string
	min := math.MaxInt32

	for _, s := range strs {
		if len(s) < min {
			min = len(s)
			shortestStr = s
		}
	}

	// 边界处理
	if shortestStr == "" {
		return ""
	}

	for min > 0 {
		flag := true
		for _, s := range strs {
			if !strings.HasPrefix(s, shortestStr[:min]) {
				flag = false
				break
			}
		}
		if flag {
			commonPrefix = shortestStr[:min]
			break
		}
		min--
	}

	return commonPrefix

}
