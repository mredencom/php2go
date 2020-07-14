package php

import "os"

// https://www.php.net/manual/zh/function.mkdir.php
// mkdir()
func Mkdir(name string, mode os.FileMode) bool {
	err := os.Mkdir(name, mode)
	return err == nil
}
