package php

import "strings"

// https://www.php.net/manual/zh/function.str-split.php
// str_split
func StrSplit(s string, sep string, n int) []string {
	return strings.SplitN(s, sep, n)
}
