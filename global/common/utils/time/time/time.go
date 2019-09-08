package time

import (
	"time"
	"strconv"
)

//获取当前时间戳
func CurrentTime() int {
	t := time.Now()
	timestamp,_ := strconv.Atoi(strconv.FormatInt(t.UTC().UnixNano(),10)[:10])
	return timestamp
}

//获取当前时间
func CurrentTimeNow() time.Time {
	return time.Now()
}

//日期赚时间戳 | 秒
func DateToTimestamp(date time.Time) int {
	d,_ := time.Parse("2019-04-18 00:00:00",date.String())
	timestamp,_ := strconv.Atoi(strconv.FormatInt(d.UTC().UnixNano(),10)[:10])
	return timestamp
}