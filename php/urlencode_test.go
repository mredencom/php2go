package php

import "testing"

func TestUrlencode(t *testing.T) {
	ss := Urlencode("https://www.baidu.com")
	t.Log(ss)
	r := Urldecode(ss)
	t.Log(r)
}
