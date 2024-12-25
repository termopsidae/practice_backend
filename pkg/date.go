package pkg

import (
	"paractice/config"
	"strconv"
	"time"
)

func GetLastDay() int64 {
	currentTime := time.Now().In(config.BeijingLoc)
	oldTime := currentTime.AddDate(0, 0, -1)
	y, m, d := oldTime.Date()
	date := int64(y*10000 + int(m)*100 + d)
	return date
}
func GetToday() int64 {
	currentTime := time.Now().In(config.BeijingLoc)
	y, m, d := currentTime.Date()
	date := int64(y*10000 + int(m)*100 + d)
	return date
}
func GetStringToday() string {
	currentTime := time.Now().In(config.BeijingLoc)
	y, m, d := currentTime.Date()
	date := int64(y*10000 + int(m)*100 + d)
	dateStr := strconv.FormatInt(date, 10)
	s := dateStr[0:4] + "-" + dateStr[4:6] + "-" + dateStr[6:]
	return s
}
func GetStringLastToday() string {
	currentTime := time.Now().In(config.BeijingLoc)
	oldTime := currentTime.AddDate(0, 0, -1)
	y, m, d := oldTime.Date()
	date := int64(y*10000 + int(m)*100 + d)
	dateStr := strconv.FormatInt(date, 10)
	s := dateStr[0:4] + "-" + dateStr[4:6] + "-" + dateStr[6:]
	return s
}
func GetDurationToNextDay() int64 {
	now := time.Now().In(config.BeijingLoc)
	// 获取明天凌晨的时间
	tomorrow := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())

	duration := tomorrow.Unix() - now.Unix()
	return duration
}
