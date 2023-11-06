package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Parse(input string) (CronSchedule, string, error) {
	components := strings.Split(input, " ")
	if len(components) < 6 {
		return CronSchedule{}, "", errors.New("invalid cron expression")
	}

	minute, minuteErr := possibleFormat.Parse(minuteRange, components[0])
	hour, hourErr := possibleFormat.Parse(hourRange, components[1])
	dayOfMonth, dayOfMonthErr := possibleFormat.Parse(dayOfMonthRange, components[2])
	month, monthErr := possibleFormat.Parse(monthRange, components[3])
	dayOfWeek, dayOfWeekErr := possibleFormat.Parse(dayOfWeekRange, components[4])
	command := components[5]

	for _, err := range []error{minuteErr, hourErr, dayOfMonthErr, monthErr, dayOfWeekErr} {
		if err != nil {
			return CronSchedule{}, "", err
		}
	}

	return CronSchedule{Minute: minute, Hour: hour, DayOfMonth: dayOfMonth, Month: month, DayOfWeek: dayOfWeek}, command, nil
}

func (f CronFormat) Parse(exRange expressionRange, input string) ([]int, error) {
	if input == "*" {
		return starExpression(exRange, input)
	} else if f.commaRegex.MatchString(input) {
		return f.commaExpression(exRange, input)
	} else if f.stepRegex.MatchString(input) {
		return f.stepExpression(exRange, input)
	} else if f.rangeRegex.MatchString(input) {
		return rangeExpression(exRange, input)
	} else if v, err := strconv.Atoi(input); err == nil {
		return []int{v}, nil
	}
	return nil, errors.New("unknown expression")
}

func starExpression(exRange expressionRange, expression string) ([]int, error) {
	if expression != "*" {
		return nil, errors.New("input must be *")
	}
	return enumerateResultWithRange(exRange.getMin(), exRange.getMax(), 1), nil
}

// rangeExpression parses expressions of the format "x-y"
func rangeExpression(exRange expressionRange, expression string) ([]int, error) {
	start, end, err := getRangeValues(exRange, expression)
	if err != nil {
		return nil, err
	}
	return enumerateResultWithRange(start, end, 1), nil
}

// stepExpression parses expressions of the format "x/y" e.g. */y, a-b/y
func (f CronFormat) stepExpression(exRange expressionRange, expression string) ([]int, error) {
	components := strings.Split(expression, "/")
	if len(components) > 2 {
		return nil, errors.New("invalid format for range, accepted format: x-y")
	}

	exp := components[0]
	step, err := validateStepValue(exRange, components[1])
	if err != nil {
		return nil, fmt.Errorf("invalid format for step value: %v", components[1])
	}

	if exp == "*" {
		return enumerateResultWithRange(exRange.getMin(), exRange.getMax(), step), nil
	} else if f.rangeRegex.MatchString(exp) {
		start, end, err := getRangeValues(exRange, exp)
		if err != nil {
			return nil, err
		}
		return enumerateResultWithRange(start, end, step), nil
	} else if v, err := strconv.Atoi(exp); err == nil {
		return enumerateResultWithRange(v, exRange.getMax(), step), nil
	}

	return nil, nil
}

// commaExpression parses expressions of the format "x,y" e.g. a,b or a,b,c or a,b-c/y
func (f CronFormat) commaExpression(exRange expressionRange, expression string) ([]int, error) {
	components := strings.Split(expression, ",")

	var res []int

	for _, c := range components {
		if f.stepRegex.MatchString(c) {
			arr, err := f.stepExpression(exRange, c)
			if err != nil {
				return nil, err
			}
			res = append(res, arr...)
		} else if f.rangeRegex.MatchString(c) {
			arr, err := rangeExpression(exRange, c)
			if err != nil {
				return nil, err
			}
			res = append(res, arr...)
		} else if c == "*" {
			return enumerateResultWithRange(exRange.getMin(), exRange.getMax(), 1), nil
		} else if v, err := strconv.Atoi(c); err == nil {
			res = append(res, v)
		} else {
			return nil, fmt.Errorf("invalid format for comma component: %v", c)
		}
	}
	return removeDuplicates(res), nil
}
