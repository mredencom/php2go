package php

import "unicode/utf8"

// https://www.php.net/manual/zh/function.mb-strlen.php
// mb_strlen
func MbStrlen(s string) int {
	return utf8.RuneCountInString(s)
}
