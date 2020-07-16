package php

import "strings"

// https://www.php.net/manual/zh/function.stripos.php
// stripos
// 获取字符串首次出现的位置 不区分大小写
func Stripos(haystack, needle string, offset int) int {
	l := len(haystack)
	if l == 0 || offset > l || -offset > l {
		return -1
	}
	haystack = haystack[offset:]
	if offset < 0 {
		offset += l
	}
	pos := strings.Index(strings.ToLower(haystack), strings.ToLower(needle))
	if pos == -1 {
		return -1
	}
	return pos + offset
}
