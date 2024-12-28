package jira

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const LogFilePath = "time_logs.json"

type DailyLog struct {
	Date  string  `json:"date"`
	Hours float64 `json:"hours"`
}

type TimeLogHistory struct {
	Logs []DailyLog `json:"logs"`
}

// LoadHistory reads the logs from the JSON file
func LoadHistory() (*TimeLogHistory, error) {
	file, err := os.Open(filepath.Clean(LogFilePath))
	if os.IsNotExist(err) {
		return &TimeLogHistory{Logs: []DailyLog{}}, nil
	}
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var history TimeLogHistory
	err = json.NewDecoder(file).Decode(&history)
	if err != nil {
		return nil, err
	}
	return &history, nil
}

// SaveHistory writes the logs to the JSON file
func SaveHistory(history *TimeLogHistory) error {
	file, err := os.Create(filepath.Clean(LogFilePath))
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	return json.NewEncoder(file).Encode(history)
}

// AddOrUpdateLog adds hours to the log for a specific date
func AddOrUpdateLog(history *TimeLogHistory, date string, hours float64) {
	for i, log := range history.Logs {
		if log.Date == date {
			history.Logs[i].Hours += hours
			return
		}
	}

	history.Logs = append(history.Logs, DailyLog{
		Date:  date,
		Hours: hours,
	})
}
