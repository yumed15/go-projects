package scheduler

import (
	"cron-parser/parser"
	"time"
)

func calculateNextRun(schedule parser.CronSchedule, timeNow time.Time) Schedule {
	givenTime := Schedule{timeNow.Minute(), timeNow.Hour(),
		timeNow.Day(), int(timeNow.Month())}

	minute, nextHour := getNext(schedule.Minute, givenTime.minute)

	hour := timeNow.Hour()
	nextDayOfMonth := false
	if nextHour {
		hour, nextDayOfMonth = getNext(schedule.Hour, givenTime.hour)
	}

	dayOfMonth := timeNow.Day()
	nextMonth := false
	if nextDayOfMonth {
		dayOfMonth, nextMonth = getNext(schedule.DayOfMonth, givenTime.day)
	}

	month := int(timeNow.Month())
	if nextMonth {
		month, _ = getNext(schedule.Month, givenTime.month)
	}

	return Schedule{minute, hour, dayOfMonth, month}

}

func getNext(arr []int, val int) (int, bool) {
	for i := 0; i < len(arr); i++ {
		if arr[i] > val {
			return arr[i], false
		}
	}
	return arr[0], true
}
