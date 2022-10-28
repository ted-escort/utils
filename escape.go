package utils

import (
	"bytes"
	"encoding/json"
)

// DisableEscape 去除json中的转义字符
func DisableEscape(data any) (string, error) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	if err := jsonEncoder.Encode(data); err != nil {
		return "", err
	}
	return bf.String(), nil
}
