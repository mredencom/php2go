package php

import "math"

// https://www.php.net/manual/zh/function.atan.php
// 反正切
// x float64 度数
func Atan(x float64) float64 {
	return math.Atan(x)
}
