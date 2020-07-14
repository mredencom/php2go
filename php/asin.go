package php

import "math"

// https://www.php.net/manual/zh/function.asin.php
// 反正弦函数
// x float64 度数
func Asin(x float64) float64 {
	return math.Asin(x)
}
