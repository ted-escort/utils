package utils

func InitSlice(len int) []map[string]any {
	data := make([]map[string]any, len)
	for k := range data {
		data[k] = make(map[string]any)
	}
	return data
}

func InitMap(len int) map[string]any {
	data := make(map[string]any, len)
	return data
}
