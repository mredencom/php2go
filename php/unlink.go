package php

import "os"

// https://www.php.net/manual/zh/function.unlink.php
// unlink()
func Unlink(filename string) bool {
	err := os.Remove(filename)
	return err == nil
}
