package util

import (
	"time"
)

func UnixToTime(unix int64) *time.Time {
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return nil
	}
	date, err := time.ParseInLocation("2006-01-02 15:04:05", time.Unix(unix, 0).Format("2006-01-02 15:04:05"), location)
	if err != nil {
		return nil
	}
	return &date
}

func DateStringToTime(dateString string) *time.Time {
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return nil
	}
	date, err := time.ParseInLocation("2006-01-02 15:04:05", dateString, location)
	if err != nil {
		return nil
	}
	return &date
}
