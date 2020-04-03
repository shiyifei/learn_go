package config

/**
	jwt sha256加密算法的盐值
 */

var JwtShaSecret []byte

func init() {
	JwtShaSecret = []byte("JWT!20200403secret")
}

