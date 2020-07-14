package php

import "path/filepath"

// https://www.php.net/manual/zh/function.glob.php
// glob()
func Glob(pattern string) []string {
	match, err := filepath.Glob(pattern)
	if err != nil {
		return []string{}
	}
	return match
}
