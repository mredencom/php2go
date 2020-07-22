package php

import "testing"

func TestArrayPop(t *testing.T) {
	ss := []interface{}{
		"hello",
		111,
		222,
		[]string{"你好"},
	}
	t.Log(ArrayPop(ss))
}
