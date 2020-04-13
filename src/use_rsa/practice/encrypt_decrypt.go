package practice

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"crypto/rand"
	"encoding/pem"
	"fmt"
)

//RSA对称性加密
//公钥加密、私钥解密
var privateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQCe8zGb4UAMg2A63pH+/W145hHvYQPJlkX6OfzJ1215htCI6Pyh
2TdHRrDqVU6wP609ao9tLxRsbbXrajBGXiq2ijRX7AKrsVdhYi2J+B2q/CrsH5CD
Ka16YCVPPwf/oZDz/hxrcjZjhOoSIZupY3/xzOBTTjcVcvWbTxGw0wOm6wIDAQAB
AoGABrVg7KFPILgSwalPJCHyEt4y95VyoXl0LqFv59ztw+lKt9yNfQ875Ag5w0oi
bhHh7+ulbghEpmbi/LKYov+qccTQMCz4PW1g85LrUYI5PaGKQfsTAWldQeV/mxCk
mimCk8bahoWPX4i2fnyFdCCn7f3kL8RqRp4NXu2En2gJkPECQQDL3QZrRBpxuE8L
vgMPNew+II3XtiMzsXc/EwHpAT2hY/pOXt0pvtGfAU2d1JSzmHlBfqPkhr2S0obE
PpdsXyG3AkEAx5mt8rsDflY8vRYU7Xao0+Smt+9ujMhvtzzS9W62VCUU8xc0UG+x
umgxofSOedkoaR7k2jqFYYbC1CrwPyAUbQJBALle2R9gZctSFE5REOcb2R0E7PVg
oNG4ZP3tgqckga3nAwuQJvp2kJVM0g7Z5f0If/mV9eEuw+JlnDWF1JquRjECQQCi
ZrT0eRsnkO0MgEn4yAInnbPUlphhLbhP48pVbYYmQqGgBHJJPAfkfmBbwMqn83uA
xGU59kGOD4K39FPTWLulAkAngU3Yv8vYmZKcYXuc/TZjxa0sMuRVroWO6ciW81so
+sFpf0SM9Ysgf/nKtux7juJABCfF1ffDQdKwederSMOc
-----END RSA PRIVATE KEY-----`)

//声明公钥
//公钥可以公开给所有人使用，可以用作加密，可以用作验签
var publicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCe8zGb4UAMg2A63pH+/W145hHv
YQPJlkX6OfzJ1215htCI6Pyh2TdHRrDqVU6wP609ao9tLxRsbbXrajBGXiq2ijRX
7AKrsVdhYi2J+B2q/CrsH5CDKa16YCVPPwf/oZDz/hxrcjZjhOoSIZupY3/xzOBT
TjcVcvWbTxGw0wOm6wIDAQAB
-----END PUBLIC KEY-----`)



func RSAEncrypt(originData []byte) []byte {
	block, ok := pem.Decode(publicKey)
	fmt.Println("ok:", ok)
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)

	pub := pubInterface.(*rsa.PublicKey)

	bits, _ := rsa.EncryptPKCS1v15(rand.Reader, pub, originData)
	return bits
}

func RSADecrypt(input []byte) []byte {
	//通过私钥解密
	block, _ := pem.Decode(privateKey)
	//解析私钥
	pri, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	//解密
	bits,_ := rsa.DecryptPKCS1v15(rand.Reader, pri, input)

	return bits
}


//RSA公钥私钥产生
func GenRsaKey(bits int) (publicKeyStr, privateKeyStr string, err error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	bufferPrivate := new(bytes.Buffer)
	err = pem.Encode(bufferPrivate, block)
	if err != nil {
		return
	}
	privateKeyStr = bufferPrivate.String()
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	bufferPublic := new(bytes.Buffer)
	err = pem.Encode(bufferPublic, block)
	if err != nil {
		return
	}
	publicKeyStr = bufferPublic.String()
	fmt.Println("-------------公钥----------------")
	fmt.Println(publicKeyStr)
	fmt.Println("--------------私钥---------------")
	fmt.Println(privateKeyStr)
	return

}


func Test() {
	//生成公钥私钥
	GenRsaKey(512)

	//加密
	cipher := RSAEncrypt([]byte("hello, shiyifei"))
	fmt.Println(cipher)

	//解密
	plain := RSADecrypt(cipher)
	fmt.Println(string(plain))
}