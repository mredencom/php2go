package php

import "os"

// https://www.php.net/manual/zh/function.stat.php
// stat()
func Stat(filename string) (os.FileInfo, error) {
	return os.Stat(filename)
}
