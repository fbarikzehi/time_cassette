package utils

import (
	"strconv"
	"time"
)

func UnixToDateTime(unixStr string) (time.Time, error) {
	i, err := strconv.ParseInt(unixStr, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(i, 0), nil
}
