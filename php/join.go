package php

import "strings"

// https://www.php.net/manual/zh/function.join.php
// join
func Join(sep string, ss []string) string {
	return strings.Join(ss, sep)
}
