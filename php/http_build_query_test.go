package php

import "testing"

func TestHttpBuildQuery(t *testing.T) {
	ss := make(map[string]string)
	ss["name"] = "King"
	ss["age"] = "28"
	ss["color"] = "red green"
	t.Log(HttpBuildQuery(ss))
}
