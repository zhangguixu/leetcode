package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isAnagram_v2("anagram", "nagaram"))
	fmt.Println(isAnagram_v2("rat", "cat"))
	fmt.Println(isAnagram_v2("ab", "b"))
}

/*
	hashtable 
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countMap := make(map[byte]int, len(s))

	for i := 0; i < len(s); i++ {
		countMap[s[i]]++
		countMap[t[i]]--
	}

	for _, count := range countMap {
		if count != 0 {
			return false
		}
	}

	return true
}

/*
	因为题目假设只包含小写字母，因此使用数组替代hashtable也是没有问题的
*/
func isAnagram_v1(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}

	counter := [26]int{}
	for i := 0; i < len(s); i++ {
		counter[s[i]-'a']++
		counter[t[i]-'a']--
	}

	for _, count := range counter {
		if count != 0 {
			return false
		}
	}

	return true
}

/*
	数组技巧：假设数组如果是有序的，那么是否可以简化实际的成本

	可以拓宽思路

	时间复杂度：O(nlog(n))
	这里其实也有额外的空间成本，得看go语言string和byte之间的类型转换成本了
*/
type List []byte

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	return l[i] < l[j]
}

func isAnagram_v2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sList := []byte(s)
	tList := []byte(t)
	sort.Sort(List(sList))
	sort.Sort(List(tList))

	return string(sList) == string(tList)
}