package utils

import (
	"crypto/rand"
	"math/big"
)

// GetRandomString 返回一个指定长度的随机字符串，仅包含小写字母和数字
func GetRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	charsetLength := big.NewInt(int64(len(charset)))
	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		randomString[i] = charset[index.Int64()]
	}
	return string(randomString), nil
}
