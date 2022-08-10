package stringlib

import (
	"strconv"
	"strings"
	"time"
)

// Example : FirstOfMonth("02-2022") or FirstOfMonth("2-2022")
// You also can use time.Now().Format("01-2006") as argument
func FirstOfMonth(dt string) (time.Time, error) {
	dtList := strings.Split(dt, "-")
	intYear, err := strconv.Atoi(dtList[1])
	if err != nil {
		return time.Time{}, err
	}
	intMonth, err := strconv.Atoi(dtList[0])
	if err != nil {
		return time.Time{}, err
	}
	firstOfMonth := time.Date(intYear, time.Month(intMonth), 1, 0, 0, 0, 0, time.UTC)
	return firstOfMonth, err
}

// Example : EndOfMonth("02-2022") or EndOfMonth("2-2022")
// You also can use time.Now().Format("01-2006") as argument
func EndOfMonth(dt string) (time.Time, error) {
	firstOfMonth, err := FirstOfMonth(dt)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	if err != nil {
		return time.Time{}, err
	}
	return lastOfMonth, err
}
