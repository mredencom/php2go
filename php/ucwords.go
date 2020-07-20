package php

import "strings"

// https://www.php.net/manual/zh/function.ucwords.php
// ucwords
func Ucwords(s string) string {
	return strings.Title(s)
}
