package php

import "strings"

// https://www.php.net/manual/zh/function.str-word-count.php
// str_word_count()
func StrWordCount(s string) []string {
	return strings.Fields(s)
}
