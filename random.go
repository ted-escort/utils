package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// RandomCode 生成随机字符串
func RandomCode(length int) string {
	numerical := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numerical)
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	var builder strings.Builder
	for i := 0; i < length; i++ {
		_, _ = fmt.Fprintf(&builder, "%d", numerical[rand.Intn(r)])
	}
	return builder.String()
}
