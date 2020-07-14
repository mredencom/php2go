package php

import "testing"

func TestParseUrl(t *testing.T) {
	ss, _ := ParseUrl("http://www.baidu.com/hello?us=11")

	t.Log(ss.Scheme)
	t.Log(ss.Host)
	t.Log(ss.ForceQuery)
	t.Log(ss.User.String())
}
