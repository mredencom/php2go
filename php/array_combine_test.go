package php

import "testing"

func TestArrayCombine(t *testing.T) {
	k := ArraySlice{
		"a",
		"b",
		"c",
		1,
		1.3,
	}
	v := ArraySlice{
		"a",
		"b",
		"c",
		1,
		1.3,
	}
	r, err := ArrayCombine(k, v)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log(r)
	}
}
