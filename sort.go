package utils

import (
	"math"
	"strconv"
)

const (
	MaxInt = math.MaxInt32 // int64(^uint32(0) >> 1) // 2147483647 // 最大整数
)

// SortFormat 排序字段格式化
func SortFormat(sort int64) string {
	if sort <= 0 || sort >= MaxInt {
		return ""
	} else {
		return strconv.Itoa(int(sort))
	}
}
