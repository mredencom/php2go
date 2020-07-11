package php

import "testing"

func TestIsFloat(t *testing.T) {
	t.Log(IsFloat(1.30))
	t.Log(IsFloat(1))
	t.Log(IsFloat("4444"))
	t.Log(IsFloat(true))
	t.Log(IsFloat(false))
}
