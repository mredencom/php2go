package php

import "math"

// https://www.php.net/manual/zh/function.atanh.php
// 反双曲正切
// x float64 度数
func Atanh(x float64) float64 {
	return math.Atanh(x)
}
