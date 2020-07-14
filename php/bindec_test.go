package php

import "testing"

func TestBindec(t *testing.T) {
	t.Log(Bindec("110001001001100010"))                   //201314
	t.Log(Bindec("111111101111000011011111010010010001")) // 68435178641
}
