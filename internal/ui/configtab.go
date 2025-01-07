package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Toqn/jira-lens/internal/config"
)

// CreateConfigTab creates a tab for configuration settings
func CreateConfigTab(prefs *config.Preferences, windowCanvas fyne.Canvas) *fyne.Container {
	jiraURLEntry := widget.NewEntry()
	jiraURLEntry.SetPlaceHolder("Enter Jira URL")
	jiraURLEntry.SetText(prefs.GetJiraURL())

	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Enter Username")
	usernameEntry.SetText(prefs.GetUsername())

	apiTokenEntry := widget.NewPasswordEntry()
	apiTokenEntry.SetPlaceHolder("Enter API Token")
	apiTokenEntry.SetText(prefs.GetAPIToken())

	saveButton := widget.NewButton("Save", func() {
		prefs.SetJiraURL(jiraURLEntry.Text)
		prefs.SetUsername(usernameEntry.Text)
		prefs.SetAPIToken(apiTokenEntry.Text)

		widget.ShowPopUp(widget.NewLabel("Settings saved"), windowCanvas)
	})

	content := container.NewVBox(
		widget.NewLabel("Jira URL"),
		jiraURLEntry,
		widget.NewLabel("Jira User Name"),
		usernameEntry,
		widget.NewLabel("Jira API Token"),
		apiTokenEntry,
		saveButton,
	)

	return content
}
