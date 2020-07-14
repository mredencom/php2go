package php

import "syscall"

// https://www.php.net/manual/zh/function.is-readable.php
// is_readable()
func IsReadable(filename string) bool {
	_, err := syscall.Open(filename, syscall.O_RDONLY, 0)
	return err == nil
}
