package utils

// RemoveComma 去除末尾逗号
func RemoveComma(str string) string {
	if str == "" {
		return str
	}
	if str[len(str)-1:] == "," { // 如果最后一位是“,”，则去掉
		str = str[0 : len(str)-1]
	}
	return str
}
