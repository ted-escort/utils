package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	str := "123456"
	password := md5.New()
	password.Write([]byte(str))
	fmt.Println(hex.EncodeToString(password.Sum(nil)))
}
