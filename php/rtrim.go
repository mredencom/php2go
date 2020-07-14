package php

import "strings"

// https://www.php.net/manual/zh/function.rtrim.php
// rtrim()
func Rtrim(s string) string {
	return strings.TrimRight(s, " ")
}
