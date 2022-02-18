package go_example

import (
	"encoding/base64"
	"fmt"
)

func base64Encode(bty []byte) string {
	return base64.StdEncoding.EncodeToString(bty)
}

func base64Decode(str string) []byte {
	val, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(fmt.Sprintf("base64 decode error: %v", err))
	}
	return val
}
