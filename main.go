package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/Toqn/jira-lens/internal/config"
	"log"

	"github.com/Toqn/jira-lens/internal/ui"
)

const JiraPng = "internal/assets/jira.png"

func main() {
	jiralens := app.NewWithID("test")
	appWindow := jiralens.NewWindow("Jira Lens")
	appWindow.Resize(fyne.NewSize(800, 600))
	userPreferences := config.NewPreferences(jiralens)

	jiraIcon, err := fyne.LoadResourceFromPath(JiraPng)
	if err != nil {
		log.Fatal(err)
	}

	configTab := ui.CreateConfigTab(userPreferences, appWindow.Canvas())
	jiraTab, _ := ui.CreateJiraTab()

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Jira", jiraIcon, jiraTab),
		container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), configTab),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	appWindow.SetContent(tabs)
	appWindow.ShowAndRun()
}
