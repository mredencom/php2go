package php

import "testing"

func TestIsInteger(t *testing.T) {
	t.Log(IsInteger(1.3))
	t.Log(IsInteger(1))
	t.Log(IsInteger("4444"))
	t.Log(IsInteger(true))
	t.Log(IsInteger(false))
}
