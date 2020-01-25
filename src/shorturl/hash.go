package shorturl

import (
	"crypto/md5"
	"fmt"
)

func GenerateHash(src string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(src)))
}
