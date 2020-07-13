package php

import "testing"

func TestStrtoupper(t *testing.T) {
	t.Log(Strtoupper("hello World!"))
	t.Log(Strtoupper("你好s"))
	t.Log(Strtoupper("(*^▽^*)"))
}
