package go_example

import (
	"reflect"
	"testing"
)

func Test_byte2str(t *testing.T) {
	type args struct {
		bty []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1",
			args{
				bty: []byte("test1"),
			},
			"test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := byte2str(tt.args.bty); got != tt.want {
				t.Errorf("byte2str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_float2str(t *testing.T) {
	type args struct {
		val float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				val: 3.14,
			},
			"3.14",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := float2str(tt.args.val); got != tt.want {
				t.Errorf("float2str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_int2str(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				val: 6,
			},
			"6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := int2str(tt.args.val); got != tt.want {
				t.Errorf("int2str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_json2map(t *testing.T) {
	type args struct {
		jsn string
	}
	want := make(map[string]float64)
	want["test"] = 123

	tests := []struct {
		name string
		args args
		want map[string]float64
	}{
		{
			"test1",
			args{
				jsn: "{\"test\": 123}",
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := json2map(tt.args.jsn); got["test"] != tt.want["test"] {
				t.Errorf("json2map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_map2json(t *testing.T) {
	type args struct {
		mp map[string]interface{}
	}
	mp := make(map[string]interface{})
	mp["test"] = 123
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				mp,
			},
			"{\"test\":123}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := map2json(tt.args.mp); got != tt.want {
				t.Errorf("map2json() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_str2byte(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"test1",
			args{
				str: "123",
			},
			[]byte("123"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := str2byte(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("str2byte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_str2float(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"test1",
			args{
				str: "3.14",
			},
			3.14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := str2float(tt.args.str); got != tt.want {
				t.Errorf("str2float() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_str2int(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				str: "123",
			},
			123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := str2int(tt.args.str); got != tt.want {
				t.Errorf("str2int() = %v, want %v", got, tt.want)
			}
		})
	}
}
