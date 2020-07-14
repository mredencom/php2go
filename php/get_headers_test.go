package php

import "testing"

func TestGetHeaders(t *testing.T) {
	h := GetHeaders("https://www.baidu.com")
	t.Log(h)
}
