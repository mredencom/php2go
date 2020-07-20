package php

import "strings"

// https://www.php.net/manual/zh/function.strstr.php
// strstr
func Strstr(haystack, needle string) string {
	if len(needle) == 0 {
		return ""
	}
	idx := strings.Index(haystack, needle)
	if idx == -1 {
		return ""
	}
	return haystack[idx+len([]byte(needle))-1:]
}
