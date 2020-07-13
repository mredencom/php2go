package php

import "testing"

func TestDate(t *testing.T) {
	t.Log(Date("02/01/2006 15:04:05 PM", 1594646722))
	t.Log(Date("2006-01-02 15:04:05", 1594646722))
	t.Log(Date("2006-01-02", 1594646722))
	t.Log(Date("15:04:05", 1594646722))
}
