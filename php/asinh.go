package php

import "math"

// https://www.php.net/manual/zh/function.asinh.php
// 反双曲正弦函数
// x float64 度数
func Asinh(x float64) float64 {
	return math.Asinh(x)
}
