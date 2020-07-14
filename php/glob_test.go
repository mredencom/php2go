package php

import "testing"

func TestGlob(t *testing.T) {
	t.Log(Glob("*.go"))
}
