package utils

import "time"

func TimePassed(t time.Time) bool {
	return time.Now().Sub(t) > 0
}

func FormattedTimeWithDate(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func FormattedTime(t time.Time) string {
	return t.Format("15:04:05")
}

func FormattedDate(t time.Time) string {
	return t.Format("2006-01-02")
}
