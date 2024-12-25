package qr

import (
	qrcode "github.com/skip2/go-qrcode"
)

func GenQRCode(s string) ([]byte, error) {
	// 生成二维码
	png, err := qrcode.Encode(s, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return png, nil
}
