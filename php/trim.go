package php

import "strings"

// https://www.php.net/manual/zh/function.trim.php
// trim()
func Trim(s string) string {
	return strings.Trim(s," ")
}
