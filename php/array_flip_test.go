package php

import "testing"

func TestArrayFlip(t *testing.T) {
	r := make(map[interface{}]interface{})
	r["a"] = 1
	r["b"] = 2
	r["c"] = 3
	t.Log(ArrayFlip(r))
}
