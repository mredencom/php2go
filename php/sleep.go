package php

import "time"

// https://www.php.net/manual/zh/function.sleep.php
// sleep()
func Sleep(s int64) {
	time.Sleep(time.Duration(s) * time.Second)
}
