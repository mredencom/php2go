package php

import "math"

// https://www.php.net/manual/zh/function.pow.php
// pow()
func Pow(base, exp float64) float64 {
	return math.Pow(base, exp)
}
