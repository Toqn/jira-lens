package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"jira-lens/pkg/config"
	"log"

	"jira-lens/pkg/ui"
)

const JIRA_PNG = "pkg/assets/jira.png"

func main() {
	jiralens := app.NewWithID("test")
	appWindow := jiralens.NewWindow("Jira Lens")
	appWindow.Resize(fyne.NewSize(800, 600))
	userPreferences := config.NewPreferences(jiralens)

	jiraIcon, err := fyne.LoadResourceFromPath(JIRA_PNG)
	if err != nil {
		log.Fatal(err)
	}

	configTab := ui.CreateConfigTab(userPreferences, appWindow.Canvas())
	jiraTab := ui.CreateJiraTab()

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Jira", jiraIcon, jiraTab),
		container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), configTab),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	appWindow.SetContent(tabs)
	appWindow.ShowAndRun()
}
