package php

import "math"

// https://www.php.net/manual/zh/function.acosh.php
// 反双曲余弦函数
// x float64 度数
func Acosh(x float64) float64 {
	return math.Acosh(x)
}
