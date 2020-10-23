package php

import "strings"

// https://www.php.net/manual/zh/function.explode.php
// explode()
func Explode(delimiter, s string) []string {
	return strings.Split(s, delimiter)
}
