package config

import "fyne.io/fyne/v2"

// PreferenceKeys
const (
	JiraURLKey  = "jira_url"
	UsernameKey = "username"
	APITokenKey = "api_token"
)

// Preferences wraps the Fyne preferences API
type Preferences struct {
	prefs fyne.Preferences
}

// NewPreferences creates a new Preferences wrapper
func NewPreferences(app fyne.App) *Preferences {
	return &Preferences{prefs: app.Preferences()}
}

// GetJiraURL retrieves the Jira URL
func (p *Preferences) GetJiraURL() string {
	return p.prefs.String(JiraURLKey)
}

// SetJiraURL saves the Jira URL
func (p *Preferences) SetJiraURL(value string) {
	p.prefs.SetString(JiraURLKey, value)
}

// GetUsername retrieves the username
func (p *Preferences) GetUsername() string {
	return p.prefs.String(UsernameKey)
}

// SetUsername saves the username
func (p *Preferences) SetUsername(value string) {
	p.prefs.SetString(UsernameKey, value)
}

// GetAPIToken retrieves the API token
func (p *Preferences) GetAPIToken() string {
	return p.prefs.String(APITokenKey)
}

// SetAPIToken saves the API token
func (p *Preferences) SetAPIToken(value string) {
	p.prefs.SetString(APITokenKey, value)
}
