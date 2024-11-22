package md5_util

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(data []byte) string {
	sum := md5.Sum(data)
	return hex.EncodeToString(sum[:])
}
