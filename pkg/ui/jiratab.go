package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"jira-lens/pkg/jira"
	"strconv"
	"strings"
	"time"
)

var recentTickets []string
var favoriteTickets map[string]bool

func init() {
	recentTickets = []string{}
	favoriteTickets = make(map[string]bool)
}

// CreateJiraTab creates a Jira tab with a functional recent tickets list
func CreateJiraTab() (*fyne.Container, error) {
	// Input Section
	ticketEntry := widget.NewEntry()
	ticketEntry.SetPlaceHolder("Enter ticket ID (e.g., INF-1337)")

	dateEntry := widget.NewEntry()
	dateEntry.SetPlaceHolder("YYYY-MM-DD")
	dateEntry.SetText(time.Now().Format("2006-01-02"))

	hoursEntry := widget.NewEntry()
	hoursEntry.SetPlaceHolder("Hours")
	hoursEntry.SetText("0")

	minutesEntry := widget.NewEntry()
	minutesEntry.SetPlaceHolder("Minutes")
	minutesEntry.SetText("0")

	descriptionEntry := widget.NewMultiLineEntry()
	descriptionEntry.SetPlaceHolder("Enter task description")

	// Recent Tickets List
	recentList := widget.NewList(
		func() int {
			return len(recentTickets)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("") // Template for list items
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(recentTickets[id])
		},
	)

	// On selection, populate the Ticket ID field
	recentList.OnSelected = func(id widget.ListItemID) {
		if id >= 0 && id < len(recentTickets) {
			ticketEntry.SetText(recentTickets[id])
		}
	}

	// Log Time Button
	logTimeButton := widget.NewButton(
		"Log Time",
		func() {
			// Validate inputs
			hours, err := strconv.Atoi(hoursEntry.Text)
			if err != nil {
				dialog.ShowError(err, fyne.CurrentApp().Driver().AllWindows()[0])
				return
			}
			minutes, err := strconv.Atoi(minutesEntry.Text)
			if err != nil {
				dialog.ShowError(err, fyne.CurrentApp().Driver().AllWindows()[0])
				return
			}

			// Log time
			jira.LogTime(
				ticketEntry.Text,
				dateEntry.Text,
				hours,
				minutes,
				descriptionEntry.Text,
			)

			// Update recent tickets and refresh the list
			updateRecentTickets(ticketEntry.Text)
			recentList.Refresh()
		},
	)

	// Layout
	content := container.NewVBox(
		widget.NewLabel("Log Time"),
		widget.NewLabel("Ticket ID"),
		ticketEntry,
		widget.NewLabel("Date"),
		dateEntry,
		widget.NewLabel("Hours & Minutes"),
		container.NewGridWithColumns(2, hoursEntry, minutesEntry),
		widget.NewLabel("Description"),
		descriptionEntry,
		logTimeButton,
		widget.NewLabelWithStyle("Recent Tickets", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		recentList,
	)

	return content, nil
}

// updateRecentTickets ensures the recent tickets list stays unique and within limits
func updateRecentTickets(ticketID string) {
	if ticketID == "" {
		return
	}

	// Standardize the ticket ID (e.g., convert to uppercase)
	ticketID = strings.ToUpper(ticketID)

	// Remove the ticket if it already exists (move it to the front)
	for i, t := range recentTickets {
		if t == ticketID {
			// Remove the ticket from its current position
			recentTickets = append(recentTickets[:i], recentTickets[i+1:]...)
			break
		}
	}

	// Add the ticket to the front of the list
	recentTickets = append([]string{ticketID}, recentTickets...)

	// Enforce a maximum limit of 10 tickets
	if len(recentTickets) > 10 {
		recentTickets = recentTickets[:10]
	}
}
