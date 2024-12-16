package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"jira-lens/pkg/jira"
	"strconv"
	"time"
)

// CreateJiraTab creates a tab for configuration settings
func CreateJiraTab() (*fyne.Container, error) {
	ticketEntry := widget.NewEntry()
	ticketEntry.SetPlaceHolder("Enter ticketEntry id (e.g., INF-1337)")

	dateEntry := widget.NewEntry()
	dateEntry.SetPlaceHolder("YYYY-MM-DD")
	dateEntry.SetText(time.Now().Format("2006-01-02"))

	hoursEntry := widget.NewEntry()
	hoursEntry.SetPlaceHolder("Hours")
	hoursEntry.SetText("0")
	hours, err := strconv.Atoi(hoursEntry.Text)
	if err != nil {
		return nil, err
	}

	minutesEntry := widget.NewEntry()
	minutesEntry.SetPlaceHolder("Minutes")
	minutesEntry.SetText("0")
	minutes, err := strconv.Atoi(minutesEntry.Text)
	if err != nil {
		return nil, err
	}

	descriptionEntry := widget.NewMultiLineEntry()
	descriptionEntry.SetPlaceHolder("Description")
	descriptionEntry.SetText("")

	logTimeButton := widget.NewButton(
		"log",
		func() {
			jira.LogTime(
				ticketEntry.Text,
				dateEntry.Text,
				hours,
				minutes,
				descriptionEntry.Text,
			)
		},
	)

	content := container.NewVBox(
		widget.NewLabel("Ticket ID"),
		ticketEntry,
		widget.NewLabel("Date"),
		dateEntry,
		widget.NewLabel("Time"),
		hoursEntry,
		minutesEntry,
		logTimeButton,
	)

	return content, nil
}
