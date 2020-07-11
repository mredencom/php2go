package php

import "testing"

func TestIsDouble(t *testing.T) {
	t.Log(IsDouble(1.30)) // true
	t.Log(IsDouble(1)) // false
	t.Log(IsDouble("4444")) // false
	t.Log(IsDouble(true)) // false
	t.Log(IsDouble(false)) // false
}
