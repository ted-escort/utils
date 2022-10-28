package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	password := md5.New()
	password.Write([]byte(str))
	return hex.EncodeToString(password.Sum(nil))
}
