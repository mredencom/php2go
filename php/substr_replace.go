package php

import "strings"

// https://www.php.net/manual/zh/function.str-replace.php
// str_replace
func StrReplace(search, old, new string, count int) string {
	return strings.Replace(search, old, new, count)
}
