package shorturl

import (
	"crypto/md5"
	"fmt"
)

func GenerateHash(dest string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(dest)))
}
