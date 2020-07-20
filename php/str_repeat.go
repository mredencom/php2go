package php

import "strings"

// https://www.php.net/manual/zh/function.str-repeat.php
// str-repeat
func StrRepeat(s string, multiplier int) string {
	return strings.Repeat(s, multiplier)
}
