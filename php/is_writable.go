package php

import "syscall"

// https://www.php.net/manual/zh/function.is-writable.php
// is_writable()
func IsWritable(filename string) bool {
	_, err := syscall.Open(filename, syscall.O_WRONLY, 0)
	return err == nil
}
