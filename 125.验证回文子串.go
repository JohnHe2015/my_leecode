/*
•	题目编号：125
  - 题目名称：Valid Palindrome
  - 题目描述：给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
    空字符串被定义为有效的回文串。
*/
package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(validPalindrome("A man, a plan, a canal: Panama")) // 输出: true
	fmt.Println(validPalindrome("race a car"))
}

func validPalindrome(s string) bool {
	// 将所有字母和数字字符提取出来，并转为小写
	var filtered []rune
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			filtered = append(filtered, unicode.ToLower(ch))
		}
	}

	left, right := 0, len(filtered)-1
	for left < right {
		if filtered[left] != filtered[right] {
			return false
		}
		left++
		right--
	}
	return true
}
