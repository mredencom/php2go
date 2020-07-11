package php

import "testing"

func TestBase64Encode(t *testing.T) {
	t.Log(Base64Encode("hello world"))
}

func TestBase64Decode(t *testing.T) {
	ss, _ := Base64Decode("aGVsbG8gd29ybGQ=")
	t.Log(ss)
}
