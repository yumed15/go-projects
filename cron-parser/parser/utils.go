package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func validateValue(exRange expressionRange, value string) (int, error) {
	res, err := strconv.Atoi(strings.TrimSpace(value))
	if err != nil {
		return 0, err
	}
	if res < exRange.getMin() || res > exRange.getMax() {
		return 0, errors.New("invalid value")
	}
	return res, nil
}

func validateStepValue(exRange expressionRange, value string) (int, error) {
	res, err := strconv.Atoi(strings.TrimSpace(value))
	if err != nil {
		return 0, err
	}
	return exRange.formatStepValue(res), nil
}

func enumerateResultWithRange(min int, max int, step int) []int {
	result := make([]int, 0)
	for i := min; i <= max; i += step {
		result = append(result, i)
	}
	return result
}

func getRangeValues(exRange expressionRange, expression string) (int, int, error) {
	components := strings.Split(expression, "-")
	if len(components) > 2 {
		return 0, 0, errors.New("invalid format for range, accepted format: x-y")
	}

	start, errStart := validateValue(exRange, components[0])
	end, errEnd := validateValue(exRange, components[1])
	if errStart != nil || errEnd != nil || start > end {
		return 0, 0, fmt.Errorf("invalid range provided: %v-%v", start, end)
	}

	return start, end, nil
}

func FormatOutput(schedule CronSchedule, command string) string {
	return fmt.Sprintf("%-14s%s\n", "minute", formatResult(schedule.Minute)) +
		fmt.Sprintf("%-14s%s\n", "hour", formatResult(schedule.Hour)) +
		fmt.Sprintf("%-14s%s\n", "day of month", formatResult(schedule.DayOfMonth)) +
		fmt.Sprintf("%-14s%s\n", "month", formatResult(schedule.Month)) +
		fmt.Sprintf("%-14s%s\n", "day of week", formatResult(schedule.DayOfWeek)) +
		fmt.Sprintf("%-14s%s\n", "command", command)
}

func formatResult(arr []int) string {
	var res string

	for _, el := range arr {
		res += strconv.Itoa(el) + " "
	}

	return res[:len(res)-1]
}

func removeDuplicates[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func replaceAlphaWithNumeric(cronExpr string) string {
	monthMap := map[string]string{
		"JAN": "1", "FEB": "2", "MAR": "3", "APR": "4",
		"MAY": "5", "JUN": "6", "JUL": "7", "AUG": "8",
		"SEP": "9", "OCT": "10", "NOV": "11", "DEC": "12",
	}

	dayOfWeekMap := map[string]string{
		"SUN": "0", "MON": "1", "TUE": "2", "WED": "3",
		"THU": "4", "FRI": "5", "SAT": "6",
	}

	parts := strings.Fields(cronExpr)
	for i, part := range parts {
		if val, ok := monthMap[part]; ok {
			parts[i] = val
		} else if val, ok := dayOfWeekMap[part]; ok {
			parts[i] = val
		}
	}
	return strings.Join(parts, " ")
}
