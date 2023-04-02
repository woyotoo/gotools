package util

import (
	"testing"
	"time"
)

// cd gotools
// go test -v util/datetime_test.go util/datetime.go

func Test_HumanTime(t *testing.T) {
	// var timestamp int64 = 1625882650
	timestamp := time.Now().Unix()

	type testCase struct {
		Name      string
		Timestamp int64
		Expected  string
	}
	testCases := []testCase{
		{"Test - 刚刚", timestamp - 5, "刚刚"},
		{"Test - 1 分钟前", timestamp - 100, "1 分钟前"},
		{"Test - 60 分钟前", timestamp - 3600, "60 分钟前"},
		{"Test - 3 天前", timestamp - 3600*24*3, "3 天前"},
		{"Test - 2 周前", timestamp - 3600*24*20, "2 周前"},
		{"Test - 1 月前", timestamp - 3600*24*35, "1 月前"},
		{"Test - 2 月前", timestamp - 3600*24*66, "2 月前"},
		{"Test - 3 月前", timestamp - 3600*24*100, "3 月前"},
		{"Test - 1 年前", timestamp - 3600*24*364, "1 年前"},
		{"Test - 1 年前", timestamp - 3600*24*366, "1 年前"},
		{"Test - 1 年前", timestamp - 3600*24*400, "1 年前"},
	}
	for _, item := range testCases {
		result := HumanTime(item.Timestamp)
		if result != item.Expected {
			t.Errorf("Test_HumanTime failed. Name[%v], Input[%v], Got[%v], expected[%v]", item.Name, item.Timestamp, result, item.Expected)
		}
	}
}

func Test_Str2HumanTime(t *testing.T) {
	t.Log("---> Test_Str2HumanTime")

	// data := "2021-02-06 06:25:56"
	data := "2021-08-06T06:25:42Z"

	str := Str2HumanTime(data)

	t.Log("---> Test_Str2HumanTime", str)
}

func Test_UnixTime2DateTime(t *testing.T) {
	t.Log("---> Test_UnixTime2DateTime")

	data := 1612563956

	result := UnixTime2DateTime(int64(data))
	t.Log("---> UnixTime2DateTime unixtime", data)
	t.Log("---> UnixTime2DateTime datetime", result)
}

func Test_Str2UnixTime(t *testing.T) {
	t.Log("---> Str2UnixTime")

	data := "2021-02-06 06:25:56"

	result := Str2UnixTime(data)

	t.Log("---> Str2UnixTime datetime", data)
	t.Log("---> Str2UnixTime unixtime", result)
}

func Test_TimeAgo(t *testing.T) {
	type args struct {
		t time.Time
	}

	myTime := time.Now().Add(-72 * time.Hour)

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"小于1分钟",
			args{time.Now().Add(-30 * time.Second)},
			"约 1 分钟前",
		},
		{
			"约2小时前",
			args{time.Now().Add(-100 * time.Minute)},
			"约 2 小时前",
		},
		{
			"约1天前",
			args{time.Now().Add(-25 * time.Hour)},
			"约 1 天前",
		},
		{
			"同年的从月份开始",
			args{myTime},
			myTime.Format("01-02 15:04"),
		},
		{
			"2016-02-02",
			args{time.Date(2016, 2, 2, 0, 0, 0, 0, time.Local)},
			"2016-02-02",
		},
	}
	for _, info := range tests {
		t.Run(info.name, func(t *testing.T) {
			got := TimeAgo(info.args.t)

			if got != info.want {
				t.Errorf("Test_TimeAgo failed: got %v, want %v", got, info.want)
			}
		})
	}
}
