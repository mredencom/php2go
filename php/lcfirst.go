package php

import "unicode"

// https://www.php.net/manual/zh/function.lcfirst.php
// lcfirst()
func Lcfirst(s string) string {
	for _, v := range s {
		u := string(unicode.ToLower(v))
		return u + s[len(u):]
	}
	return ""
}
