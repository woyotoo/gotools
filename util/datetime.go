package util

import (
	"fmt"
	"strings"
	"time"
)

// UnixTime time()
func UnixTime() int64 {
	return time.Now().Unix()
}

// 时间戳转成 datetime 字符串 YY-mm-dd
func UnixTime2Date(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02")
}

// 时间戳转成 datetime 字符串 YY-mm-dd HH:ii:ss
func UnixTime2DateTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

// 将时间字符串转成可读性友好的时间字符串
func Str2HumanTime(datetime string) string {

	layout := "2006-01-02 15:04:05"
	index := strings.Index(datetime, "T")
	if index == 10 {
		layout = "2006-01-02T15:04:05Z"
	}
	localTime, _ := time.ParseInLocation(layout, datetime, time.Local)
	res := localTime.Unix()
	fmt.Println(" datetime ", datetime)
	fmt.Println(" localTime ", localTime)
	fmt.Println(" index ", index)

	return HumanTime(res)
}

// UTC时间字符串转成时间戳
func Str2UnixTime(datetime string) int64 {
	layout := "2006-01-02 15:04:05"
	index := strings.Index(datetime, "T")
	if index == 10 {
		layout = "2006-01-02T15:04:05Z"
	}
	localTime, _ := time.ParseInLocation(layout, datetime, time.Local)

	return localTime.Unix()
}

// UTC时间字符串转成时间
func Str2Datetime(datetime string) string {
	layout := "2006-01-02 15:04:05"
	index := strings.Index(datetime, "T")
	if index == 10 {
		layout = "2006-01-02T15:04:05Z"
	}
	localTime, _ := time.ParseInLocation(layout, datetime, time.Local)

	return localTime.Format("2006-01-02 15:04:05")
}

// 获取可读性友好的时间字符串
func HumanTime(timestamp int64) string {
	const minute = 60 // 60 秒
	const hour = 3600 // 3600 秒
	const day = hour * 24
	const week = day * 7
	const month = day * 30
	const year = month * 12
	var now = time.Now().Unix() // 获取当前时间戳
	var diff = now - timestamp  // 给的日期与当前时间戳的差值
	// startTime := time.Now()
	// durations := time.Since(startTime)
	// time.Sleep(time.Second * 2)
	// fmt.Println("离现在过去了：", durations.Seconds())

	var s string
	if diff > year {
		s = fmt.Sprintf("%.0f 年前", float64(diff/year))
	} else if diff > month {
		s = fmt.Sprintf("%.0f 月前", float64(diff/month))
	} else if diff > week {
		s = fmt.Sprintf("%.0f 周前", float64(diff/week))
	} else if diff > day {
		s = fmt.Sprintf("%.0f 天前", float64(diff/day))
	} else if diff > hour {
		s = fmt.Sprintf("%.0f 小时前", float64(diff/hour))
	} else if diff > minute {
		s = fmt.Sprintf("%.0f 分钟前", float64(diff/minute))
	} else if diff > 1000 {
		s = fmt.Sprintf("%.0f 秒前", float64(diff/1000))
	} else {
		s = "刚刚"
	}

	return s
}

// TimeAgo
func TimeAgo(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)
	hours := diff.Hours()
	if hours < 1.0 {
		return fmt.Sprintf("约 %.0f 分钟前", diff.Minutes())
	}
	if hours < 24.0 {
		return fmt.Sprintf("约 %.0f 小时前", hours)
	}
	if hours < 72.0 {
		return fmt.Sprintf("约 %.0f 天前", hours/24.0)
	}
	// 同一年的显示月份
	if now.Year() == t.Year() {
		return t.Format("01-02 15:04")
	}

	return t.Format("2006-01-02")
}
