package php

import "os"

// https://www.php.net/manual/zh/function.chown.php
// chown
func Chown(name string, uid, gid int) bool {
	err := os.Chown(name, uid, gid)
	if err != nil {
		return false
	}
	return true
}
