package php

import "testing"

func TestSha1File(t *testing.T) {
	s1, _ := Sha1File("http://wxsnsdythumb.wxs.qq.com/109/20204/snsvideodownload?m=9c702d45d56ee7452dd50f217ecc863c&filekey=30340201010420301e02016d0402534804109c702d45d56ee7452dd50f217ecc863c0203013a50040d00000004627466730000000131&hy=SH&storeid=32303230303732303130353030323030303536316435313336666664393335323738343530393030303030303664&bizid=1023")
	s2, _ := Sha1File("http://copy-writing.oss-cn-shanghai.aliyuncs.com/copywriting/inside/images/9c702d45d56ee7452dd50f217ecc863c.png")
	s3, _ := Sha1File("http://copy-writing.oss-cn-shanghai.aliyuncs.com/copywriting/outerside/images/OSGbxMRkWP7iYBobLJV7uQHXbDQiVOn3uTDzbIgx.mp4")
	t.Log("s1:", s1)
	t.Log("s2:", s2)
	t.Log("s3:", s3)
}
