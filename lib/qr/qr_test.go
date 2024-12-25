package qr

import (
	"os"
	"testing"
)

func TestQr(t *testing.T) {
	s := "https://www.baidu.com"
	data, err := GenQRCode(s)
	if err != nil {
		t.Logf("GenQRCode err:%s", err)
	}
	t.Logf("data:%v", data)
	f, err := os.Create("qr.png")
	if err != nil {
		t.Logf("os.Create err:%s", err)
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		t.Logf("f.Write err:%s", err)
	}
}
