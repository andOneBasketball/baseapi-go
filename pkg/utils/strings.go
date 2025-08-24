package utils

import "strings"

// MaskMiddle 遮掩字符串中间一半数据
func MaskMiddle(s string) string {
	n := len(s)
	if n <= 2 {
		return strings.Repeat("*", n)
	}

	// 取前 1/4 和后 1/4 的长度，至少保留一个字符
	left := n / 4
	right := n / 4
	if left == 0 {
		left = 1
	}
	if right == 0 {
		right = 1
	}

	// 需要遮掩的中间部分长度
	maskLen := n - left - right
	if maskLen <= 0 {
		maskLen = 1
	}

	return s[:left] + strings.Repeat("*", maskLen) + s[n-right:]
}
