package php

import "strings"

// https://www.php.net/manual/zh/function.ltrim.php
// ltrim()
func Ltrim(s string) string {
	return strings.TrimLeft(s, " ")
}
