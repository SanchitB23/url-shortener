package utils

import (
	"github.com/segmentio/ksuid"
	"os"
	"strconv"
)

func CreateShortUUID() string {
	var maxLen int64
	maxLen, err := strconv.ParseInt(os.Getenv("URL_UUID_LEN"), 10, 64)
	if err != nil || maxLen < 6 {
		maxLen = 7
	}
	return ksuid.New().String()[:maxLen]
}
