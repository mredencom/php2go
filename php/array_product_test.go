package php

import "testing"

func TestArrayProduct(t *testing.T) {
	sr := []float64{
		1, 2, 1.2,
	}

	t.Log(ArrayProduct(sr))
}
