package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi(" 	-42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("+122"))
	fmt.Println(myAtoi("     +004500"))
	fmt.Println(myAtoi("   -04f"))
	fmt.Println(myAtoi("   +0 123"))
	fmt.Println(myAtoi("-5-"))
	fmt.Println(myAtoi("-5+"))
}

// 解法的关键在于ascii编码的数值， http://ascii.911cha.com/
const (
	Space = 32 // 空格
	Tab   = 9  // tab键
	Minus = 45 // -号
	Plus  = 43
)

func isNumber(b byte) bool {
	if b >= 48 && b <= 57 {
		return true
	}
	return false
}

func isLegal(b byte) bool {
	if isNumber(b) || b == Space || b == Tab || b == Minus || b == Plus {
		return true
	}
	return false
}

func myAtoi(str string) int {
	var num float64
	total := len(str) - 1
	idx := 0

	isMinus := false

	for i, _ := range str {
		idx = i
		b := str[i]
		if !isLegal(b) {
			break
		}
		// 处理空格和tab
		if b == Space || b == Tab || b == Minus || b == Plus {
			if i > 0 && (str[i-1] != Space && str[i-1] != Tab) {
				break
			}
			if b == Minus {
				isMinus = true
			}
			continue
		}

		num += math.Pow10(total-i) * float64(b-48)
	}

	// 处理 4193 with words，必须回溯到空格之前
	if num > 0 && idx <= total {
		for idx >= 0 {
			b := str[idx]
			if isLegal(b) {
				// 回退
				if b == Space || b == Tab || b == Plus || b == Minus {
					idx--
				}
				break
			}
			idx--
		}
		num = num / math.Pow10(total-idx)
	}

	// 处理负号
	if isMinus {
		num = -num
	}

	// 处理边界
	if num < -math.Pow(2, 31) {
		num = -math.Pow(2, 31)
	}

	if num > math.Pow(2, 31)-1 {
		num = math.Pow(2, 31) - 1
	}

	return int(num)
}
