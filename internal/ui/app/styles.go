package app

import "charm.land/lipgloss/v2"

var (
	colorBg          = lipgloss.Color("#151326")
	colorPanel       = lipgloss.Color("#1b1833")
	colorPanelSoft   = lipgloss.Color("#211e3d")
	colorBorder      = lipgloss.Color("#4b4673")
	colorBorderFocus = lipgloss.Color("#d7d0ff")

	colorText  = lipgloss.Color("#f4f1ff")
	colorMuted = lipgloss.Color("#8e88b7")

	colorGreen  = lipgloss.Color("#7ee787")
	colorAmber  = lipgloss.Color("#ffb86c")
	colorRed    = lipgloss.Color("#ff7b72")
	colorPurple = lipgloss.Color("#bd93f9")
	colorBlue   = lipgloss.Color("#79c0ff")

	colorCreamWhite = lipgloss.Color("#f5f0d7")
	colorCharcoal   = lipgloss.Color("#151326")
	colorLightGrey  = lipgloss.Color("#34315e")
)

var (
	cursorStyle   = lipgloss.NewStyle().Background(colorCreamWhite).Foreground(colorCharcoal).Bold(true)
	selectedStyle = lipgloss.NewStyle().Background(colorLightGrey).Foreground(colorText)

	borderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorBorder).
			Padding(0, 1)

	focusedBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(colorBorderFocus).
				Padding(0, 1)

	resourceBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(colorBorder).
				Padding(0, 1)

	dimStyle      = lipgloss.NewStyle().Foreground(colorMuted)
	moduleStyle   = lipgloss.NewStyle().Foreground(colorText).Bold(true)
	errorStyle    = lipgloss.NewStyle().Foreground(colorRed).Bold(true)
	warningStyle  = lipgloss.NewStyle().Foreground(colorAmber).Bold(true)
	successStyle  = lipgloss.NewStyle().Foreground(colorGreen).Bold(true)
	infoBarStyle  = lipgloss.NewStyle().Foreground(colorText).Bold(true)
	helpKeyStyle  = lipgloss.NewStyle().Background(colorPurple).Foreground(colorCreamWhite).Bold(true).Padding(0, 1)
	helpDescStyle = lipgloss.NewStyle().Foreground(colorMuted)

	treePrefixDefaultStyle = lipgloss.NewStyle().Foreground(colorBorder)
	treePrefixCurrentStyle = lipgloss.NewStyle().Foreground(colorBorderFocus).Bold(true)
)
