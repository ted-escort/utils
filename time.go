package utils

import (
	"bytes"
	"math"
	"strconv"
	"time"
)

const (
	TimeFormatYmdHis = "2006-01-02 15:04:05"
	TimeFormatYmd    = "2006-01-02"
)

// Timestamp 获取当前时间戳
func Timestamp() int64 {
	return time.Now().Unix()
}

// TimeToDateFull 时间戳转年月日时分秒
func TimeToDateFull(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(TimeFormatYmdHis)
}

// TimeToDate 时间戳转年月日
func TimeToDate(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(TimeFormatYmd)
}

// DateFull 时间戳转年月日时分秒
func DateFull() string {
	return time.Unix(Timestamp(), 0).Format(TimeFormatYmdHis)
}

// Date 时间戳转年月日
func Date() string {
	return time.Unix(Timestamp(), 0).Format(TimeFormatYmd)
}

// DateFullToTime 日期（年月日时分秒）转时间戳
func DateFullToTime(date string) int64 {
	timestamp, err := time.ParseInLocation(TimeFormatYmdHis, date, time.Local)
	if err != nil {
		return 0
	}
	return timestamp.Unix()
}

// DateToTime 日期（年月日）转时间戳
func DateToTime(date string) int64 {
	timestamp, err := time.ParseInLocation(TimeFormatYmd, date, time.Local)
	if err != nil {
		return 0
	}
	return timestamp.Unix()
}

func FormatTime(atime int64) string {
	var byTime = []int64{365 * 24 * 60 * 60, 30 * 24 * 60 * 60, 24 * 60 * 60, 60 * 60, 60, 1}
	var unit = []string{"年前", "个月前", "天前", "小时前", "分钟前", "秒钟前"}
	now := time.Now().Unix()
	ct := now - atime
	if ct < 0 {
		return "刚刚"
	}
	var res string
	for i := 0; i < len(byTime); i++ {
		if ct < byTime[i] {
			continue
		}
		var temp = math.Floor(float64(ct / byTime[i]))
		ct = ct % byTime[i]
		if temp > 0 {
			var tempStr string
			tempStr = strconv.FormatFloat(temp, 'f', -1, 64)
			res = MergeString(tempStr, unit[i]) // 字符串拼接的
		}
		break // 精确到最大单位，即："2天前"这种形式，如果想要"2天12小时36分钟48秒前"这种形式，把此处break去掉，然后把字符串拼接调整下即可
	}
	return res
}

func MergeString(args ...string) string {
	buffer := bytes.Buffer{}
	for i := 0; i < len(args); i++ {
		buffer.WriteString(args[i])
	}
	return buffer.String()
}
