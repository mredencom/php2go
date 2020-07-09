package php

import "testing"

func TestArrayCountValues(t *testing.T) {
	var ss []interface{}
	ss = append(ss, "hello")
	ss = append(ss, "hello")
	ss = append(ss, "hello")
	ss = append(ss, "world")
	ss = append(ss, "world")
	ss = append(ss, "world")
	ss = append(ss, "66")
	t.Log(ArrayCountValues(ss))
}
