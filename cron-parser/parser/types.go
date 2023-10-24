package parser

import "regexp"

type CronFormat struct {
	rangeRegex *regexp.Regexp
	stepRegex  *regexp.Regexp
	commaRegex *regexp.Regexp
}

var possibleFormat = CronFormat{
	rangeRegex: regexp.MustCompile(`\d-\d`),
	stepRegex:  regexp.MustCompile(`./\d`),
	commaRegex: regexp.MustCompile(`.,.`),
}

type CronSchedule struct {
	Minute     []int
	Hour       []int
	DayOfMonth []int
	Month      []int
	DayOfWeek  []int
}
