package php

import "math"

// https://www.php.net/manual/zh/function.atan2.php
// 两个参数的反正切
// x y float64 度数
func Atan2(x, y float64) float64 {
	return math.Atan2(x, y)
}
