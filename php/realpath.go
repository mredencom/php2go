package php

import "path/filepath"

// https://www.php.net/manual/zh/function.realpath.php
// realpath()
func Realpath(filename string) string {
	s, err := filepath.Abs(filename)
	if err != nil {
		return ""
	}
	return s
}
