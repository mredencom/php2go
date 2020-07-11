package php

import "testing"

func TestArrayMap(t *testing.T) {
	var a []interface{}

	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	res := ArrayMap(func(item interface{}) interface{} {
		i := item.(int)
		return i * i
	}, a)
	t.Log(res)
}
