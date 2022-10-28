package utils

import (
	"encoding/json"
	"fmt"
)

// CopyStruct 拷贝结构体
func CopyStruct(source, target any) error {
	bytes, _ := json.Marshal(source)
	err := json.Unmarshal(bytes, target)
	if err != nil {
		fmt.Println("结构体转换失败：", err.Error())
	}
	return err
}

// StructToMap 结构体转map
func StructToMap(source any) (map[string]any, error) {
	sourceJson, err := json.Marshal(source)
	if err != nil {
		return nil, err
	}
	targetMap := InitMap(0)
	err = json.Unmarshal(sourceJson, &targetMap)
	if err != nil {
		return nil, err
	}
	return targetMap, nil
}
