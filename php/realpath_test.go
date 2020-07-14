package php

import "testing"

func TestRealpath(t *testing.T) {
	t.Log(Realpath("testdata"))
}
