package php

import "testing"

func TestArrayPush(t *testing.T) {
	as := []interface{}{
		"hello",
		"world",
		111,
	}
	ArrayPush(&as, 1, 2, 3, 4, "hello")
	t.Log(as)
}
