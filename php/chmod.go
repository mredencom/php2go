package php

import "os"

// https://www.php.net/manual/zh/function.chmod.php
// chmod
func Chmod(name string, mode os.FileMode) bool {
	err := os.Chmod(name, mode)
	if err != nil {
		return false
	}
	return true
}
