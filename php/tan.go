package php

import "math"

// https://www.php.net/manual/zh/function.tan.php
// 正切
// x float64 度数
func Tan(x float64) float64 {
	return math.Tan(x)
}
