package php

import "math"

// https://www.php.net/manual/zh/function.cos.php
// 增加余弦函数
// x float64 度数
func Cos(x float64) float64 {
	return math.Cos(x)
}
