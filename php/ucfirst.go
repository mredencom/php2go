package php

import "unicode"

// https://www.php.net/manual/zh/function.ucfirst.php
// ucfirst()
func Ucfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return u + str[len(u):]
	}
	return ""
}
