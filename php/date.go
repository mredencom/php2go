package php

import "time"

// https://www.php.net/manual/zh/function.date.php
// date()
func Date(format string, timestamp int64) string {
	return time.Unix(timestamp, 0).Format(format)
}
