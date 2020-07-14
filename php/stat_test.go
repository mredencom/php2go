package php

import "testing"

func TestStat(t *testing.T) {
	f, err := Stat("testdata/3.txt")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(f)
}
