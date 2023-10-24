package scheduler

import (
	"cron-parser/parser"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCronExpression(t *testing.T) {
	cases := []struct {
		input    string
		schedule parser.CronSchedule
		timeNow  time.Time
		expected Schedule
	}{
		{
			input: "* * * * * /cmd",
			schedule: parser.CronSchedule{
				Minute:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59},
				Hour:       []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
				DayOfMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
				Month:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				DayOfWeek:  []int{1, 2, 3, 4, 5, 6, 7}},
			timeNow:  time.Date(2009, 11, 17, 20, 4, 0, 0, time.UTC),
			expected: Schedule{minute: 5, hour: 20, day: 17, month: 11},
		},
		{
			input: "*/10 * * * * /cmd",
			schedule: parser.CronSchedule{
				Minute:     []int{0, 10, 20, 30, 40, 50},
				Hour:       []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
				DayOfMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
				Month:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				DayOfWeek:  []int{1, 2, 3, 4, 5, 6, 7}},
			timeNow:  time.Date(2009, 11, 17, 20, 4, 0, 0, time.UTC),
			expected: Schedule{minute: 10, hour: 20, day: 17, month: 11},
		},
		{
			input: "1-10/2 * * * * /cmd",
			schedule: parser.CronSchedule{
				Minute:     []int{1, 3, 5, 7, 9},
				Hour:       []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
				DayOfMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
				Month:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				DayOfWeek:  []int{1, 2, 3, 4, 5, 6, 7}},
			timeNow:  time.Date(2009, 11, 17, 23, 9, 0, 0, time.UTC),
			expected: Schedule{minute: 1, hour: 0, day: 18, month: 11},
		},
		{
			input: "*/15 0 1,15 * 1-5",
			schedule: parser.CronSchedule{
				Minute:     []int{0, 15, 30, 45},
				Hour:       []int{0},
				DayOfMonth: []int{1, 15},
				Month:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				DayOfWeek:  []int{1, 2, 3, 4, 5}},
			timeNow:  time.Date(2009, 11, 17, 23, 9, 0, 0, time.UTC),
			expected: Schedule{minute: 15, hour: 23, day: 17, month: 11},
		},
		{
			input: "*/15 0 1,15 * 1-5",
			schedule: parser.CronSchedule{
				Minute:     []int{0, 15, 30, 45},
				Hour:       []int{0},
				DayOfMonth: []int{1, 15},
				Month:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				DayOfWeek:  []int{1, 2, 3, 4, 5}},
			timeNow:  time.Date(2009, 11, 17, 23, 47, 0, 0, time.UTC),
			expected: Schedule{minute: 0, hour: 0, day: 1, month: 12},
		},
	}

	for _, test := range cases {
		t.Run(test.input, func(t *testing.T) {
			nextRun := calculateNextRun(test.schedule, test.timeNow)
			assert.Equal(t, test.expected, nextRun)
		})
	}
}
