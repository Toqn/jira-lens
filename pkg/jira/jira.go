package jira

import "fmt"

// TimeLog represents the data required for a time log
type TimeLog struct {
	Issue       string
	Date        string
	Hours       int
	Minutes     int
	Description string
}

func LogTime(issue string, date string, hours int, minutes int, description string) TimeLog {
	log := TimeLog{
		Issue:       issue,
		Date:        date,
		Hours:       hours,
		Minutes:     minutes,
		Description: description,
	}
	fmt.Print(log)
	return log
}
