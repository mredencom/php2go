package php

import "math"

// https://www.php.net/manual/zh/function.cosh.php
// 双曲余弦
// x float64 度数
func Cosh(x float64) float64 {
	return math.Cosh(x)
}
