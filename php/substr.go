package php

// https://www.php.net/manual/zh/function.substr.php
// substr()
func Substr(s string, start uint, l int) string {
	if start < 0 || l < -1 {
		return s
	}
	switch {
	case l == -1:
		return s[start:]
	case l == 0:
		return ""
	}
	end := int(start) + l
	if end > len(s) {
		end = len(s)
	}
	return s[start:end]
}
