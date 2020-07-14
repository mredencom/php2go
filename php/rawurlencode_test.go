package php

import "testing"

func TestRawurldecode(t *testing.T) {
	sr := Rawurlencode("https://www.bilibili.com")
	t.Log(sr)
	se := Rawurldecode(sr)
	t.Log(se)
}
