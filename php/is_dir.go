package php

import "os"

// https://www.php.net/manual/zh/function.is-dir.php
// is_dir()
func IsDir(name string) bool {
	f, err := os.Stat(name)
	return err == nil && f.IsDir()
}
