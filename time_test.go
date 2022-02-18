package go_example

import (
	"testing"
	"time"
)

func TestFormatDatetime(t *testing.T) {
	type args struct {
		timestamp int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				timestamp: 1645155620,
			},
			"2022-02-18 11:40:20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDatetime(tt.args.timestamp); got != tt.want {
				t.Errorf("FormatDatetime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatDatetimeByLoc(t *testing.T) {
	type args struct {
		timestamp int64
		locStr    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				timestamp: 1645155620,
				locStr:    "Asia/Shanghai",
			},
			"2022-02-18 11:40:20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDatetimeByLoc(tt.args.timestamp, tt.args.locStr); got != tt.want {
				t.Errorf("FormatDatetimeByLoc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatNow(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				format: "2006-01-02",
			},
			"2022-02-18",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatNow(tt.args.format); got != tt.want {
				t.Errorf("FormatNow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatNowByLoc(t *testing.T) {
	type args struct {
		format string
		locStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				format: "2006-01-02",
				locStr: "Asia/Shanghai",
			},
			"2022-02-18",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatNowByLoc(tt.args.format, tt.args.locStr); got != tt.want {
				t.Errorf("FormatNowByLoc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatTimestamp(t *testing.T) {
	type args struct {
		datetime string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"test1",
			args{
				datetime: "2022-02-18 11:40:20",
			},
			1645184420,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatTimestamp(tt.args.datetime); got != tt.want {
				t.Errorf("FormatTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatTimestampByLoc(t *testing.T) {
	type args struct {
		datetime string
		locStr   string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"test1",
			args{
				datetime: "2022-02-18 11:40:20",
				locStr:   "Asia/Shanghai",
			},
			1645155620,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatTimestampByLoc(tt.args.datetime, tt.args.locStr); got != tt.want {
				t.Errorf("FormatTimestampByLoc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubDatetime(t *testing.T) {
	type args struct {
		start string
		end   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				start: "2022-02-18 11:40:20",
				end:   "2022-02-18 11:40:30",
			},
			"10s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubDatetime(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("SubDatetime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddDatetime(t *testing.T) {
	type args struct {
		datetime string
		dur      time.Duration
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				datetime: "2022-02-18 11:40:20",
				dur:      time.Duration(10) * time.Second,
			},
			"2022-02-18 11:40:30",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddDatetime(tt.args.datetime, tt.args.dur); got != tt.want {
				t.Errorf("AddDatetime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddDatetimeByLoc(t *testing.T) {
	type args struct {
		datetime string
		dur      time.Duration
		locStr   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				datetime: "2022-02-18 11:40:20",
				dur:      time.Duration(10) * time.Second,
				locStr:   "Asia/Shanghai",
			},
			"2022-02-18 11:40:30",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddDatetimeByLoc(tt.args.datetime, tt.args.dur, tt.args.locStr); got != tt.want {
				t.Errorf("AddDatetimeByLoc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiffDatetime(t *testing.T) {
	type args struct {
		start string
		end   string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"test1",
			args{
				start: "2022-02-18 11:40:20",
				end:   "2022-02-18 11:40:30",
			},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiffDatetime(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("DiffDatetime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiffDatetimeByLoc(t *testing.T) {
	type args struct {
		start       string
		end         string
		startLocStr string
		endLocStr   string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"test1",
			args{
				start:       "2022-02-18 11:40:20",
				end:         "2022-02-18 11:40:30",
				startLocStr: "Asia/Shanghai",
				endLocStr:   "Asia/Shanghai",
			},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiffDatetimeByLoc(tt.args.start, tt.args.end, tt.args.startLocStr, tt.args.endLocStr); got != tt.want {
				t.Errorf("DiffDatetimeByLoc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubDatetimeByLoc(t *testing.T) {
	type args struct {
		start       string
		end         string
		startLocStr string
		endLocStr   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				start:       "2022-02-18 11:40:20",
				end:         "2022-02-18 11:40:30",
				startLocStr: "Asia/Shanghai",
				endLocStr:   "Asia/Shanghai",
			},
			"10s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubDatetimeByLoc(tt.args.start, tt.args.end, tt.args.startLocStr, tt.args.endLocStr); got != tt.want {
				t.Errorf("SubDatetimeByLoc() = %v, want %v", got, tt.want)
			}
		})
	}
}
