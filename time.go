package go_example

import "time"

// FormatNow 格式化当前时间
func FormatNow(format string) string {
	// format 2006-01-02 15:04:05
	return time.Now().Format(format)
}

// FormatTimestamp 给定时间转时间戳
func FormatTimestamp(datetime string) int64 {
	// 默认UTC，相差8小时
	times, err := time.Parse("2006-01-02 15:04:05", datetime)
	if err != nil {
		panic(err)
	}
	return times.Unix()
}

// FormatDatetime 时间戳转时间
func FormatDatetime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

// FormatNowByLoc 格式化当前时间
func FormatNowByLoc(format, locStr string) string {
	loc, err := time.LoadLocation(locStr)
	if err != nil {
		panic(err)
	}
	// format 2006-01-02 15:04:05
	return time.Now().In(loc).Format(format)
}

// FormatTimestampByLoc 给定时间转时间戳
func FormatTimestampByLoc(datetime, locStr string) int64 {
	loc, err := time.LoadLocation(locStr)
	if err != nil {
		panic(err)
	}
	times, err := time.ParseInLocation("2006-01-02 15:04:05", datetime, loc)
	if err != nil {
		panic(err)
	}
	return times.Unix()
}

// FormatDatetimeByLoc 时间戳转时间
func FormatDatetimeByLoc(timestamp int64, locStr string) string {
	loc, err := time.LoadLocation(locStr)
	if err != nil {
		panic(err)
	}
	return time.Unix(timestamp, 0).In(loc).Format("2006-01-02 15:04:05")
}

// DiffDatetime 计算时间差
func DiffDatetime(start, end string) int64 {
	// 默认UTC，相差8小时
	startTimes, err := time.Parse("2006-01-02 15:04:05", start)
	if err != nil {
		panic(err)
	}
	endTimes, err := time.Parse("2006-01-02 15:04:05", end)
	if err != nil {
		panic(err)
	}

	return endTimes.Unix() - startTimes.Unix()
}

// DiffDatetime 计算时间差
func DiffDatetimeByLoc(start, end, startLocStr, endLocStr string) int64 {
	startLoc, err := time.LoadLocation(startLocStr)
	if err != nil {
		panic(err)
	}
	endLoc, err := time.LoadLocation(endLocStr)
	if err != nil {
		panic(err)
	}

	startTimes, err := time.ParseInLocation("2006-01-02 15:04:05", start, startLoc)
	if err != nil {
		panic(err)
	}
	endTimes, err := time.ParseInLocation("2006-01-02 15:04:05", end, endLoc)
	if err != nil {
		panic(err)
	}

	return endTimes.Unix() - startTimes.Unix()
}

// SubDatetime 计算
// return 72h3m0.5s
func SubDatetime(start, end string) string {
	// 默认UTC，相差8小时
	startTimes, err := time.Parse("2006-01-02 15:04:05", start)
	if err != nil {
		panic(err)
	}
	endTimes, err := time.Parse("2006-01-02 15:04:05", end)
	if err != nil {
		panic(err)
	}

	return endTimes.Sub(startTimes).String()
}

// SubDatetimeByLoc 计算时间差
// return 72h3m0.5s
func SubDatetimeByLoc(start, end, startLocStr, endLocStr string) string {
	startLoc, err := time.LoadLocation(startLocStr)
	if err != nil {
		panic(err)
	}
	endLoc, err := time.LoadLocation(endLocStr)
	if err != nil {
		panic(err)
	}

	startTimes, err := time.ParseInLocation("2006-01-02 15:04:05", start, startLoc)
	if err != nil {
		panic(err)
	}
	endTimes, err := time.ParseInLocation("2006-01-02 15:04:05", end, endLoc)
	if err != nil {
		panic(err)
	}
	return endTimes.Sub(startTimes).String()
}

// AddDatetime 给定时间增加指定时间
func AddDatetime(datetime string, dur time.Duration) string {
	// 默认UTC，相差8小时
	times, err := time.Parse("2006-01-02 15:04:05", datetime)
	if err != nil {
		panic(err)
	}
	return times.Add(dur).Format("2006-01-02 15:04:05")
}

// AddDatetimeByLoc 给定时间增加指定时间
func AddDatetimeByLoc(datetime string, dur time.Duration, locStr string) string {
	loc, err := time.LoadLocation(locStr)
	if err != nil {
		panic(err)
	}

	times, err := time.ParseInLocation("2006-01-02 15:04:05", datetime, loc)
	if err != nil {
		panic(err)
	}
	return times.Add(dur).Format("2006-01-02 15:04:05")
}
