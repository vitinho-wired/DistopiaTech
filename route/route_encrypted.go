package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"

	cryptoz "github.com/vitinho-wired/DistopiaTech/mygo/src/crypto"
)

var privatekey *rsa.PrivateKey
var err error
var pubkey *rsa.PublicKey

func init() {
	privatekey, err = rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Println(err.Error())

	}
	pubkey = &privatekey.PublicKey
}

func encryptrsa(msg string) string {
	if privatekey == nil {
		return "chave priva vazia"
	}

	msgencrypted, errorr := rsa.EncryptOAEP(sha256.New(), rand.Reader, pubkey, []byte(msg), nil)
	if errorr != nil {
		fmt.Println(errorr.Error())
	}
	msgstrencrypted := string(msgencrypted[:])
	return cryptoz.Encoding64(msgstrencrypted)
}

func decryptrsa(msg string) string {
	if pubkey == nil {
		return "chave publica vazia"
	}
	msgdecode64 := cryptoz.Decoding64(msg)
	msgdecrypted, error1 := rsa.DecryptOAEP(sha256.New(), rand.Reader, privatekey, []byte(msgdecode64), nil)
	if error1 != nil {
		fmt.Println(error1.Error())
	}
	msgdecryptedstr := string(msgdecrypted[:])
	return msgdecryptedstr

}

func handlerencrypt(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Println("metodo invalido")
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bodystring := string(body[:])
	bodystring = encryptrsa(bodystring)
	fmt.Println(bodystring)

}
func handlerdecrypt(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Println("metodo invalido")
		return
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bodystring := string(body[:])
	bodystring = decryptrsa(bodystring)
	fmt.Println(bodystring)
}

func main() {

	http.HandleFunc("/encrypt", handlerencrypt)
	http.HandleFunc("/decrypt", handlerdecrypt)
	fmt.Println("server listening on port 8071:")
	err := http.ListenAndServe(":8071", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}
