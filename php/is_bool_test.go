package php

import "testing"

func TestIsBool(t *testing.T) {
	t.Log(IsBool("nihao "))
	t.Log(IsBool(false))
	t.Log(IsBool(true))
	t.Log(IsBool(1))
}
