package utils

import (
	"os"
)

// GetFileByteArr 获取文件
func GetFileByteArr(fileName string) ([]byte, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}
