package go_example

import "testing"

func Test_md5(t *testing.T) {
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
			"5a105e8b9d40e1329780d62ea2265d8a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := md5(tt.args.bty); got != tt.want {
				t.Errorf("md5() = %v, want %v", got, tt.want)
			}
		})
	}
}
