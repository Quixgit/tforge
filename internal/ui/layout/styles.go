package layout

import (
	"charm.land/lipgloss/v2"

	"github.com/quix/tforge/internal/ui/theme"
)

type Styles struct {
	App       lipgloss.Style
	Header    lipgloss.Style
	Sidebar   lipgloss.Style
	Main      lipgloss.Style
	StatusBar lipgloss.Style

	MenuItem lipgloss.Style
	Selected lipgloss.Style

	Title lipgloss.Style
	Muted lipgloss.Style

	Panel lipgloss.Style
}

func NewStyles(t theme.Theme) Styles {
	return Styles{
		App: lipgloss.NewStyle().
			Background(t.Bg).
			Foreground(t.Text),

		Header: lipgloss.NewStyle().
			Background(t.Bg).
			Foreground(t.Text).
			Bold(true).
			Padding(0, 2),

		Sidebar: lipgloss.NewStyle().
			Background(t.BgPanel).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(t.BorderSubtle).
			Padding(1, 1),

		Main: lipgloss.NewStyle().
			Background(t.Bg).
			Padding(1, 2),

		StatusBar: lipgloss.NewStyle().
			Background(t.Bg).
			Foreground(t.TextDimmed).
			Padding(0, 2),

		MenuItem: lipgloss.NewStyle().
			Foreground(t.TextMuted).
			Padding(0, 1),

		Selected: lipgloss.NewStyle().
			Background(t.SelectionBg).
			Foreground(t.SelectionFg).
			Bold(true).
			Padding(0, 1),

		Title: lipgloss.NewStyle().
			Foreground(t.Accent).
			Bold(true),

		Muted: lipgloss.NewStyle().
			Foreground(t.TextMuted),

		Panel: lipgloss.NewStyle().
			Background(t.BgPanel).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(t.BorderSubtle).
			Padding(1, 2),
	}
}
