package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// CreateJiraTab creates a tab for configuration settings
func CreateJiraTab() *fyne.Container {
	ticketEntry := widget.NewEntry()
	ticketEntry.SetPlaceHolder("Enter ticketEntry id (e.g., INF-1337)")

	dateEntry := widget.NewEntry()
	dateEntry.SetPlaceHolder("Enter date")

	timeEntry := widget.NewEntry()
	timeEntry.SetPlaceHolder("Enter time spent")

	logTimeButton := widget.NewButton("log", func() {
		fmt.Printf("Logging %s on ticket %s for date %s", ticketEntry.Text, timeEntry.Text, dateEntry.Text)
	})

	content := container.NewVBox(
		widget.NewLabel("Ticket ID"),
		ticketEntry,
		widget.NewLabel("Date"),
		dateEntry,
		widget.NewLabel("Time"),
		timeEntry,
		logTimeButton,
	)

	return content
}
