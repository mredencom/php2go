package php

import "testing"

func TestRmdir(t *testing.T) {
	f := Rmdir("testdata/mk_dir")
	if !f {
		t.Error("失败")
	} else {
		t.Log("成功")
	}
}
