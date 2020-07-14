package php

import "testing"

func TestCeil(t *testing.T) {
	t.Log(Ceil(1 / 2))
	t.Log(Ceil(3 / 2))
}
