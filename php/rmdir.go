package php

import "os"

// https://www.php.net/manual/zh/function.rmdir.php
// rmdir()
func Rmdir(path string) bool {
	err := os.RemoveAll(path)
	return err == nil
}
