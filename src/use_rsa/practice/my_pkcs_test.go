package practice

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

type MyType int64

const (
	PKCS1 MyType = iota
	PKCS8
)

//默认客户端，pkcs8私钥格式，pem编码
func NewDefault(privateKey, publicKey string) (Cipher, error) {
	blockPri, _ := pem.Decode([]byte(privateKey))
	if blockPri == nil {
		return nil, errors.New("private key error")
	}

	blockPub, _ := pem.Decode([]byte(publicKey))
	if blockPub == nil {
		return nil, errors.New("public key error")
	}

	return New(blockPri.Bytes, blockPub.Bytes, PKCS8)
}

func New(privateKey, publicKey []byte, privateKeyType MyType) (Cipher, error) {
	priKey, err := genPriKey(privateKey, privateKeyType)
	if err != nil {
		return nil, err
	}
	pubKey, err := genPubKey(publicKey)
	if err != nil {
		return nil, err
	}
	return &pkcsClient{privateKey: priKey, publicKey: pubKey}, nil
}

func genPubKey(publicKey []byte) (*rsa.PublicKey, error) {
	pub, err := x509.ParsePKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	return pub.(*rsa.PublicKey), nil
}

func genPriKey(privateKey []byte, privateKeyType MyType) (*rsa.PrivateKey, error) {
	var priKey *rsa.PrivateKey
	var err error
	switch privateKeyType {
	case PKCS1:
		{
			priKey, err = x509.ParsePKCS1PrivateKey([]byte(privateKey))
			if err != nil {
				return nil, err
			}
		}
	case PKCS8:
		{
			prkI, err := x509.ParsePKCS8PrivateKey([]byte(privateKey))
			if err != nil {
				return nil, err
			}
			priKey = prkI.(*rsa.PrivateKey)
		}
	default:
		{
			return nil, errors.New("unsupport private key type")
		}
	}
	return priKey, nil
}


/*func MyTest() {
	var cipher Cipher
	client, err := NewDefault(`-----BEGIN PRIVATE KEY-----
		私钥信息
	-----END PRIVATE KEY-----`, `-----BEGIN PUBLIC KEY-----
公钥信息
-----END PUBLIC KEY-----`)
	if err != nil {
		fmt.Println(err)
	}
	cipher = client
	cp, err := cipher.Encrypt([]byte("测试加密解密"))
	if err != nil {
		fmt.Println(err)
	}
	cpStr := base64.URLEncoding.EncodeToString(cp)
	fmt.Println(cpStr)
	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		fmt.Println(err)
	}
	pp, err := cipher.Decrypt(ppBy)
	fmt.Println(string(pp))
}*/

