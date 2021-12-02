package WasteLibrary

import "strconv"

//GetDbIndexByDate
func GetDbIndexByDate(inTime string) string {
	var dbIndex string
	dbIndex = "0"

	var redisDbDate RedisDbDateType
	redisDbDate.New()

	queryTime := StringToTime(inTime)
	for ind, redisInTime := range redisDbDate.DayDates {
		redisTime := StringToTime(redisInTime)
		if redisTime.Year() == queryTime.Year() && redisTime.Month() == queryTime.Month() && redisTime.Day() == queryTime.Day() {
			dbIndex = strconv.Itoa(ind)
			break
		}
	}

	return dbIndex
}
