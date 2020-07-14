package php

import "testing"

func TestTrim(t *testing.T) {
	t.Log(Trim(" hello world"))
	t.Log(Rtrim("hello world "))
	t.Log(Ltrim(" hello world"))
}
