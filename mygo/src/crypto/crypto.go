package cryptoz

import (
	b64 "encoding/base64"
)

func Encoding64(decryptstring string) string {
	var data string = b64.StdEncoding.EncodeToString([]byte(decryptstring))
	return data
}

func Decoding64(encryptstring string) string {
	data, _ := b64.StdEncoding.DecodeString(encryptstring)
	return string(data)
}
