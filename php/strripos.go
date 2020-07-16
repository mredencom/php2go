package php

import "strings"

// https://www.php.net/manual/zh/function.strripos.php
// strripos
func Strripos(haystack, needle string, offset int) int {
	position, l := 0, len(haystack)
	if l == 0 || offset > l || -offset > l {
		return -1
	}

	if offset < 0 {
		haystack = haystack[:offset+l+1]
	} else {
		haystack = haystack[offset:]
	}
	position = strings.LastIndex(strings.ToLower(haystack), strings.ToLower(needle))
	if offset > 0 && position != -1 {
		position += offset
	}
	return position
}
