package php

import "time"

// https://www.php.net/manual/zh/function.time.php
// time()
func Time() int64 {
	return time.Now().Unix()
}
