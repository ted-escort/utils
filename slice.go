package utils

// InSlice 判断字符串是否存在数组或切片里
func InSlice(source string, slice []string) (int, bool) {
	for i, item := range slice {
		if item == source {
			return i, true
		}
	}
	return -1, false
}
