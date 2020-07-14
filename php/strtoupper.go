package php

import "strings"

// https://www.php.net/manual/zh/function.strtoupper.php
// strtoupper
func Strtoupper(s string) string {
	return strings.ToUpper(s)
}
