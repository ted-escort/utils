package utils

import (
	"bytes"
	"strconv"
	"strings"
)

// IpStringToInt ip串转int
func IpStringToInt(ipString string) int {
	ipData := strings.Split(ipString, ".")
	var ipInt int
	var pos uint = 24
	for _, ipSeg := range ipData {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}

// IpIntToString int转ip串
func IpIntToString(ipInt int) string {
	ipData := make([]string, 4)
	var ipLen = len(ipData)
	buffer := bytes.NewBufferString("")
	for i := 0; i < ipLen; i++ {
		tempInt := ipInt & 0xFF
		ipData[ipLen-i-1] = strconv.Itoa(tempInt)
		ipInt = ipInt >> 8
	}
	for i := 0; i < ipLen; i++ {
		buffer.WriteString(ipData[i])
		if i < ipLen-1 {
			buffer.WriteString(".")
		}
	}
	return buffer.String()
}
