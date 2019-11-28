package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefix_v2([]string{"flower", "flow", "flight"})) // "fl"
	fmt.Println(longestCommonPrefix_v2([]string{"dog", "racecar", "car"})) // ""
	fmt.Println(longestCommonPrefix_v2([]string{"testsss", "test", "testaas"})) //  "test"
}


/*
	这道题的解法有多种：https://leetcode-cn.com/problems/longest-common-prefix/solution/zui-chang-gong-gong-qian-zhui-by-leetcode/
	1）水平扫描法
		扫描的方向可以是从头开始，也可以从尾部开始，上述的实现就是水平扫描法，从头开始
		算法的时间复杂度：O(S)，S是所有字符串中字符数量的总和
		空间复杂度：O(1)
	2）分治法
		
	3）二分查找

	4）字典树
*/

// 水平扫描法 从头开始
func longestCommonPrefix(strs []string) string {
  var idx int
  var cur byte
  for len(strs) > 0 {
    for i := 0; i < len(strs); i++ {
      if idx == len(strs[i]) {
        return string(strs[i][:idx])
      }
      if i == 0 {
        cur = strs[i][idx]
        continue
      }
      if cur != strs[i][idx] {
        return string(strs[i][:idx])
      }
    }
    idx++
	}
	return ""
}

// 水平扫描法，从尾开始
func longestCommonPrefix_v1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix, isLCP := strs[0], true
	for len(prefix) > 0 {
		isLCP = true
		for i := 1; i < len(strs); i++ {
			if !strings.HasPrefix(strs[i], prefix) {
				isLCP = false
				prefix = prefix[:len(prefix) - 1]
				break
			}
		}
		if isLCP {
			return prefix
		}
	}
	return prefix
}

// 分治法
func longestCommonPrefix_v2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	return findLCP(strs, 0, len(strs) - 1)
}

func findLCP(strs []string, l, r int) string {
	if l == r {
		return strs[l]
	}
	m := (l + r) / 2
	leftPrefix := findLCP(strs, l, m)
	rightPrefix := findLCP(strs, m + 1, r)
	var cur string
	for i := 1; i <= len(leftPrefix) && i <= len(rightPrefix); i++ {
		if !strings.HasPrefix(rightPrefix, leftPrefix[:i]) {
			break
		}
		cur = leftPrefix[:i]
	}
	return cur
}

// 二分查找法

// 字典

