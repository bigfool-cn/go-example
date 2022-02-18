package go_example

import (
	"reflect"
	"testing"
)

func Test_base64Decode(t *testing.T) {
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
				str: "dGVzdDE=",
			},
			[]byte("test1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := base64Decode(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("base64_decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base64Encode(t *testing.T) {
	type args struct {
		bty []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				bty: []byte("test1"),
			},
			"dGVzdDE=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := base64Encode(tt.args.bty); got != tt.want {
				t.Errorf("base64_encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
