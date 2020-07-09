package php

import "math"

// https://www.php.net/manual/zh/function.acos.php
// 反余弦函数
// x float64 度数
func Acos(x float64) float64 {
	return math.Acos(x)
}
