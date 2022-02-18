package go_example

import (
	"testing"
)

func TestGenerateRSAKey(t *testing.T) {
	type args struct {
		bits int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test1",
			args{
				bits: 2048,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenerateRSAKey(tt.args.bits)
		})
	}
}

func TestRsaDecrypt(t *testing.T) {
	type args struct {
		cipherBase64Text string
		path             string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				cipherBase64Text: "yeqMWpS0zWkR+8llvs85b69YQGwBBO/bEx2fac4wkdOzia517mXHh3GWXSNgc9f+QS3sex3TBBzGBWpC6X/qLuzEm3aGdU6KE7aSb9fYbbMqHhw4sakaGLru+fQel7yodleF/xAG9arI5mAV9Q6SYEABVCfwICRGkoRuJpKfFvvYd11xupChSdpuzTFxL0olfWo7LfM7ldjSSbpZeUWBOiCS0zGjRtiPktjRNKOk0ykZzJB9kWsB0N02+Lrwrt7G9TuO/bym4534AKTXEYVLU1nQFJS4uTKvTlFU94B4SMCVzEcfeVVonTtlP8F63ZUUEwe9s6M9CNIKIJB4O14q7g==",
				path:             "D:\\Coding\\GoProject\\go-example\\private.pem",
			},
			"test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RsaDecrypt(tt.args.cipherBase64Text, tt.args.path); got != tt.want {
				t.Errorf("RsaDecrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRsaEncrypt(t *testing.T) {
	type args struct {
		plainText []byte
		path      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				plainText: []byte("test1"),
				path:      "D:\\Coding\\GoProject\\go-example\\public.pem",
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RsaEncrypt(tt.args.plainText, tt.args.path); len(got) == 0 {
				t.Errorf("RsaEncrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRsaSign(t *testing.T) {
	type args struct {
		data []byte
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"test1",
			args{
				data: []byte("test1"),
				path: "D:\\Coding\\GoProject\\go-example\\private.pem",
			},
			"kDQ7ZpqciaKPJN5CEQP3J4x5nzMtCH/7TY1u9y91vAu35uGUS9/k+QqQDbnF/y7aEDZVece0r4bywA+y6uS8x58Wph4iBFwQ7XT2yQqUWqOAvj0GNFsghxXiUxDtZ00jlu9uVRI/jhSah8AGkq6O4+ocaK/AWF5iBJd0MHdsTYmY9GFut4phznNNifQIyE/ApVgrwnvRw92YDQ0VfwtML/KHSKF81WmVvNRY0IsjRgN9Ff91SAfQPgSQUJAH7FKXLtg/oTxI6+/wujK8BpIe9+l0SIsRif4+XEnqdgAh2k2m8ZC7NOHj5z8CukXYQFpB+JR4e6B9h4yD+Uq+f1kPZQ==",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RsaSign(tt.args.data, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("RsaSign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RsaSign() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRsaSignVerify(t *testing.T) {
	type args struct {
		data       []byte
		base64Sign string
		path       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"test1",
			args{
				data:       []byte("test1"),
				base64Sign: "kDQ7ZpqciaKPJN5CEQP3J4x5nzMtCH/7TY1u9y91vAu35uGUS9/k+QqQDbnF/y7aEDZVece0r4bywA+y6uS8x58Wph4iBFwQ7XT2yQqUWqOAvj0GNFsghxXiUxDtZ00jlu9uVRI/jhSah8AGkq6O4+ocaK/AWF5iBJd0MHdsTYmY9GFut4phznNNifQIyE/ApVgrwnvRw92YDQ0VfwtML/KHSKF81WmVvNRY0IsjRgN9Ff91SAfQPgSQUJAH7FKXLtg/oTxI6+/wujK8BpIe9+l0SIsRif4+XEnqdgAh2k2m8ZC7NOHj5z8CukXYQFpB+JR4e6B9h4yD+Uq+f1kPZQ==",
				path:       "D:\\Coding\\GoProject\\go-example\\public.pem",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RsaSignVerify(tt.args.data, tt.args.base64Sign, tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("RsaSignVerify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
