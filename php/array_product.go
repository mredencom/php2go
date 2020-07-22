package php

// https://www.php.net/manual/zh/function.array-product.php
// array_product
func ArrayProduct(p []float64) float64 {
	var r float64 = 1
	for _, v := range p {
		r *= v
	}
	return r
}
