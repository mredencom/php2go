package php

import "strings"

// https://www.php.net/manual/zh/function.strpos.php
// strpos
// 获取字符串首次出现的位置 区分大小写
func Strpos(haystack, needle string, offset int) int {
	l := len(haystack)
	if l == 0 || offset > l || -offset > l {
		return -1
	}
	if offset < 0 {
		offset += l
	}
	position := strings.Index(haystack[offset:], needle)
	if position == -1 {
		return -1
	}
	return position + offset
}
