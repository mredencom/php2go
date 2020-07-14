package php

import "time"

// https://www.php.net/manual/zh/function.usleep.php
// usleep()
func Usleep(ms int64) {
	time.Sleep(time.Duration(ms) * time.Microsecond)
}
