package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(input []byte) string {
	hash := md5.Sum(input)
	return hex.EncodeToString(hash[:])
}
