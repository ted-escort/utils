package utils

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestIpStringToInt(t *testing.T) {
	ipData := strings.Split("127.0.0.1", ".")
	var ipInt int
	var pos uint = 24
	for _, ipSeg := range ipData {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	fmt.Println(ipInt)
}

func TestIpIntToString(t *testing.T) {
	ipInt := 2130706433
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
	fmt.Println(buffer.String())
}
