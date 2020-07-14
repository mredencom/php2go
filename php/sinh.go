package php

import "math"

// https://www.php.net/manual/zh/function.sinh.php
// 双曲正弦
// x float64 度数
func Sinh(x float64) float64 {
	return math.Sinh(x)
}
