
package php

import "os"

// https://www.php.net/manual/zh/function.chgrp.php
// chgrp
func Chgrp(name string, uid, gid int) bool {
	err := os.Chown(name, uid, gid)
	if err != nil {
		return false
	}
	return true
}
