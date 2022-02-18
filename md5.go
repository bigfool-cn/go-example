package go_example

import (
	md52 "crypto/md5"
	"encoding/hex"
)

func md5(bty []byte) string {
	// fmt.Sprintf("%x",md52.Sum(bty))
	h := md52.New()
	h.Write(bty)
	return hex.EncodeToString(h.Sum(nil))
}
