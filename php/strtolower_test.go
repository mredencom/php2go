package php

import "testing"

func TestStrtolower(t *testing.T) {
	t.Log(Strtolower("abCd"))
	t.Log(Strtolower("你好"))
	t.Log(Strtolower("(*^▽^*)"))
}
