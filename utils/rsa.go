package utils

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
)

func SignatureCheck(data string, signature string, key string) bool {
	hashed := sha256.Sum256([]byte(data))
	pub, _ := pem.Decode([]byte(key))
	if pub == nil {
		panic("Fail to read public key.")
	}
	interfacePubkey, err := x509.ParsePKIXPublicKey(pub.Bytes)
	if err != nil {
		panic("Fail to read public key." + err.Error())
	}
	publickey, ok := interfacePubkey.(*rsa.PublicKey)
	if !ok {
		panic("Fail to read public key.")
	}
	err = rsa.VerifyPSS(publickey, crypto.SHA256, hashed[:], []byte(signature), nil)
	if err == nil {
		return true
	} else {
		return false
	}

}
