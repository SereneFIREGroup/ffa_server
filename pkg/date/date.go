package date

import "github.com/golang-module/carbon/v2"

// TodayStart get today start timestamp
func TodayStart() int64 {
	return carbon.Parse("2020-08-05").StartOfDay().Timestamp()
}

// TodayEnd get today end timestamp
func TodayEnd() int64 {
	return carbon.Parse("2020-08-05").EndOfDay().Timestamp()
}

// DateStart get date start timestamp
func DateStart(date int64) int64 {
	y, m, d := carbon.CreateFromTimestamp(date).Date()
	return carbon.CreateFromDateTime(y, m, d, 0, 0, 0).Timestamp()
}

// DateEnd get date end timestamp
func DateEnd(date int64) int64 {
	y, m, d := carbon.CreateFromTimestamp(date).Date()
	return carbon.CreateFromDateTime(y, m, d, 23, 59, 59).Timestamp()
}
