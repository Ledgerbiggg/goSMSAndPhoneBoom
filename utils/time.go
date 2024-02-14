package utils

import "time"

// GetNowStamp 获取当前时间戳
func GetNowStamp() int64 {
	return time.Now().UnixMilli()
}
