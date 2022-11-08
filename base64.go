package utils

import (
	"encoding/base64"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

// SaveBase64 保存文件
func SaveBase64(path string, base64ImageContent string) bool {
	b, _ := regexp.MatchString(`^data:\s*image\/(\w+);base64,`, base64ImageContent)
	if !b {
		return false
	}
	re, _ := regexp.Compile(`^data:\s*image\/(\w+);base64,`)
	allData := re.FindAllSubmatch([]byte(base64ImageContent), 2)
	fileType := string(allData[0][1]) //png ，jpeg 后缀获取

	base64Str := re.ReplaceAllString(base64ImageContent, "")

	date := time.Now().Format("20060102")
	if ok := IsFileExist(path + "/" + date); !ok {
		err := os.Mkdir(path+"/"+date, 0666)
		if err != nil {
			return false
		}
	}

	var file = path + "/" + date + "/" + strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+100000) + "." + fileType
	decodeString, _ := base64.StdEncoding.DecodeString(base64Str)

	err := ioutil.WriteFile(file, decodeString, 0666)
	if err != nil {
		return false
	}
	return true
}

// IsFileExist 判断文件是否存在
func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
