package main

import (
	"crypto/rand"
	"crypto/rsa"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
)

func encodingb64(decodestring string) string {
	data := b64.StdEncoding.EncodeToString([]byte(decodestring))
	return string(data)
}

func decodingb64(encodestring string) string {
	data2, _ := b64.StdEncoding.DecodeString(encodestring)
	return string(data2)
}

func main() {
	private, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//testes
	public_json, _ := json.Marshal(private.PublicKey)
	public_string := string(public_json)
	fmt.Println(public_string)
	fmt.Println()
	public_encoding := encodingb64(public_string)
	fmt.Println(public_encoding)
	fmt.Println()
	public_decoding := decodingb64(public_encoding)
	fmt.Println(public_decoding)

	var data string = encodingb64("abc123!?$*&()'-=@~")
	fmt.Println(data)
	var data2 string = decodingb64(data)
	fmt.Println(data2)
}
