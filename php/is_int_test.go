package php

import "testing"

func TestIsInt(t *testing.T) {
	t.Log(IsInt(1.3))
	t.Log(IsInt(1))
	t.Log(IsInt("4444"))
	t.Log(IsInt(true))
	t.Log(IsInt(false))
}
