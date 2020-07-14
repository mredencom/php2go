package php

import "testing"

func TestScandir(t *testing.T) {
	fs, _ := Scandir("testdata")
	t.Log(fs)
}
