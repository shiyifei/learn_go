package config

import "use_gin/app/common"

/**
	jwt sha256加密算法的盐值
 */

var JwtShaSecret []byte

var ServerHost string

func init() {
	JwtShaSecret = []byte("JWT!20200403secret")

	ServerHost = common.GetLocalIp()
}