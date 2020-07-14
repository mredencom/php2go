package php

import "strings"

// https://www.php.net/manual/zh/function.strtolower.php
// strtolower
func Strtolower(s string) string {
	return strings.ToLower(s)
}
