package php

// https://www.php.net/manual/zh/function.empty.php
// empty
func Empty(s interface{}) bool {
	if s == "" || s == nil || s == 0 {
		return true
	}
	return false
}
