package php

import "os"

// https://www.php.net/manual/zh/function.rename.php
// rename()
func Rename(oldName, newName string) bool {
	err := os.Rename(oldName, newName)
	return err == nil
}
