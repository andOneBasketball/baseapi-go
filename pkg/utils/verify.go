package utils

import "time"

// VerifyTimestamp 验证时间戳（毫秒）
func VerifyTimestamp(timestamp int64, timeDiff ...int64) bool {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	// 允许的时间偏差（默认1min）
	allowedTimeDiff := int64(1 * 60 * 1000) // 转换为毫秒
	if len(timeDiff) > 0 {
		allowedTimeDiff = timeDiff[0]
	}
	return timestamp >= now-allowedTimeDiff && timestamp <= now+allowedTimeDiff
}
