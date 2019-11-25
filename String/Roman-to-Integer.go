package main

import "fmt"

func main() {
	fmt.Println(romanToInt_v1("LVIII")) // 58
	fmt.Println(romanToInt_v1("MCMXCIV")) // 1994
}

/*
	将罗马字转为数字，这里若干个转化规则

	Symbol       Value
	I             1
	V             5
	X             10
	L             50
	C             100
	D             500
	M             1000

	这里还有以下几种字符的组合

	IV 4 IX 9

	XL 40 XC 90

	CD  400 CM 900
*/
func romanToInt(s string) int {
	sum, n, i := 0, len(s), 0
	var num, next int
	var isLast bool
	for i < n {
		num = 0
		next = 1
		isLast = false
		if i == n-1 {
			isLast = true
		}
		switch s[i] {
		case 'I':
			num = 1
			if !isLast && s[i+1] == 'V' {
					num = 4
					next = 2
					break
			}
			if !isLast && s[i+1] == 'X' {
				num = 9
				next = 2
				break
			}
		case 'X':
			num = 10
			if !isLast && s[i+1] == 'L' {
				num = 40
				next = 2
				break
			}
			if !isLast && s[i+1] == 'C' {
				num = 90
				next = 2
				break
			}
		case 'C':
			num = 100
			if !isLast && s[i+1] == 'D' {
				num = 400
				next = 2
				break
			}
			if !isLast && s[i+1] == 'M' {
				num = 900
				next = 2
				break
			}
		case 'V':
			num = 5
		case 'L':
			num = 50
		case 'D':
			num = 500
		case 'M':
			num = 1000
		}
		i += next
		sum += num
	}
	return sum
}

// 以上的解法是很自然，当然可以实现得更加简洁，例如使用map
// 在题解中，看到有一种的思路直接使用题意给到的一个规律，如果一个数字小于它的下一个数字，那么说明是减去这个值，如果大于则是加
// 这个思路把握到了规律， 因此实现上也会比较简洁，可以在根本上减少if-else
func romanToInt_v1(s string) int {
	var sum, preNum, num int
	for i := 0; i < len(s); i++ {
		num = getValue(s[i])
		if preNum < num {
			sum -= preNum
		} else {
			sum += preNum
		}
		preNum = num
	}
	sum += preNum
	return sum
}

func getValue(b byte) int {
	var val int
	switch b {
	case 'I':
		val = 1
	case 'V':
		val = 5
	case 'X':
		val = 10
	case 'L':
		val = 50
	case 'C':
		val = 100
	case 'D':
		val = 500
	case 'M':
		val = 1000
	}
	return val
}