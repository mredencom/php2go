package php

// https://www.php.net/manual/zh/function.checkdate.php
// checkdate()
func Checkdate(m, d, y int) bool {
	if m < 1 || m > 12 || d < 1 || d > 31 || y < 1 || y > 32767 {
		return false
	}
	switch m {
	case 4, 6, 9, 11:
		if d > 30 {
			return false
		}
	case 2:
		if y%4 == 0 && (y%100 != 0 || y%400 == 0) {
			if d > 29 {
				return false
			}
		} else if d > 28 {
			return false
		}
	}
	return true
}
