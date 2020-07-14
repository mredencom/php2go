package php

import "testing"

func TestMkdir(t *testing.T) {
	f := Mkdir("testdata/mk_dir", 0777)
	if !f {
		t.Error("创建失败")
	} else {
		t.Log("创建成功")
	}

}
