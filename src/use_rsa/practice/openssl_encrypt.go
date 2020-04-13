package practice

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"reflect"

	//"io/ioutil"
	"log"
	"strings"
)


func Openssl_test() {
	//pemBytes, err := ioutil.ReadFile("sample.key.pem")
	//check(err)

	fmt.Println("strPrivateKey:", strPrivateKey)

	var pemBytes = []byte(strPrivateKey)
	fmt.Println("pemBytes:", pemBytes)

	block := firstPrivateKey(pemBytes)
	if block == nil {
		fmt.Println("No private key in input")
	}
	var err error
	derBytes := block.Bytes
	fmt.Println("derBytes:", derBytes)
	if x509.IsEncryptedPEMBlock(block) {
		fmt.Println("arrive here")
		derBytes, err = x509.DecryptPEMBlock(block, []byte("password"))
		check(err)
	}
	//var key crypto.PrivateKey
	switch block.Type {
	case "RSA PRIVATE KEY":
		key, err := x509.ParsePKCS1PrivateKey(derBytes) // or PKCS8
		check(err)
		fmt.Println(key)
	case "EC PRIVATE KEY":
		key, err := x509.ParseECPrivateKey(derBytes)
		check(err)
		fmt.Println(key)
	default:
		key, err := x509.ParsePKCS8PrivateKey(derBytes)
		check(err)
		fmt.Println("key:[", key, "]")
		fmt.Println("type of key:", reflect.TypeOf(key))

		rsaKey := key.(*rsa.PrivateKey)

		decoded, err := base64.StdEncoding.DecodeString(strPasswd)
		var password = string(decoded)
		var cipher = []byte(password)
		data, err := rsa.DecryptPKCS1v15(rand.Reader, rsaKey, cipher)
		fmt.Println("data:", string(data))
		//log.Fatal("Unsupported key type")
	}


}

func firstPrivateKey(pemBytes []byte) *pem.Block {
	var block *pem.Block
	for len(pemBytes) > 0 {
		block, pemBytes = pem.Decode(pemBytes)

		fmt.Println("block.Type:", block.Type)

		if strings.HasSuffix(block.Type, "PRIVATE KEY") {
			return block
		}
	}
	return nil
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
