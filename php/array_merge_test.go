package php

import "testing"

func TestArrayMerge(t *testing.T) {
	ss1 := make([]interface{}, 0)
	ss2 := make([]interface{}, 0)
	ss1 = append(ss1, "ss1")
	ss2 = append(ss2, "ss2")

	t.Log(ArrayMerge(ss1, ss2))
}
