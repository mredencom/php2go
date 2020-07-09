package php

import "math"

// https://www.php.net/manual/zh/function.sin.php
// 正弦函数
// x float64 度数
func Sin(x float64) float64 {
	return math.Sin(x)
}
