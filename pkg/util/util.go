package util

import (
	"strings"
	"time"

	"github.com/segmentio/ksuid"
)

func NewUUID(prefix string) string {
	if len(prefix) < 3 {
		prefix = prefix + strings.Repeat("A", 3-len(prefix))
	}
	prefix = strings.ToUpper(prefix)

	id := ksuid.New()
	return prefix + id.String()
}

func GetCurrentEpochTimeMillisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GetCurrentFormattedTimeBangkok() string {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return time.Now().In(loc).Format("2006-01-02 15:04:05")
}

func RemoveEmptyArrayString(slice []string) []string {
	var result []string
	for _, str := range slice {
		if StringIsNotEmpty(str) {
			result = append(result, str)
		}
	}
	return result
}
