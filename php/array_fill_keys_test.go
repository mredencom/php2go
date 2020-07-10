package php

import "testing"

func TestArrayFillKeys(t *testing.T) {
	var ss []interface{}

	ss = append(ss, "11")
	ss = append(ss, 22)
	ss = append(ss, 2.3)
	t.Log(ArrayFillKeys(ss, "hello world!"))

}
