package util

import (
	"time"
)

func PointerString(s string) *string     { return &s }
func PointerInt64(i int64) *int64        { return &i }
func PointerInt(i int) *int              { return &i }
func PointerBool(b bool) *bool           { return &b }
func PointerTime(t time.Time) *time.Time { return &t }

func GetAge(birthday time.Time) (age int) {
	if birthday.IsZero() {
		return 0
	}

	now := time.Now().UTC()
	age = now.Year() - birthday.Year()
	if int(now.Month()) < int(birthday.Month()) || int(now.Day()) < int(birthday.Day()) {
		age--
	}
	return age
}
