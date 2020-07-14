package php

import "os"

// https://www.php.net/manual/zh/function.is-file.php
// is_file()
func IsFile(name string) bool {
	f, err := os.Stat(name)
	return err == nil && f.Mode().IsRegular()
}
