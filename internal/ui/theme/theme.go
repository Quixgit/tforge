package theme

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

type Theme struct {
	Bg         color.Color
	BgPanel    color.Color
	BgElevated color.Color

	Border       color.Color
	BorderSubtle color.Color

	Text       color.Color
	TextMuted  color.Color
	TextDimmed color.Color

	Accent color.Color

	SelectionBg color.Color
	SelectionFg color.Color

	Green  color.Color
	Yellow color.Color
	Red    color.Color
	Blue   color.Color
}

func Dark() Theme {
	return Theme{
		Bg:           lipgloss.Color("#0a0b14"),
		BgPanel:      lipgloss.Color("#161525"),
		BgElevated:   lipgloss.Color("#1d1b31"),
		Border:       lipgloss.Color("#5c5873"),
		BorderSubtle: lipgloss.Color("#34314a"),

		Text:       lipgloss.Color("#d8d5e7"),
		TextMuted:  lipgloss.Color("#9a96ad"),
		TextDimmed: lipgloss.Color("#706c82"),

		Accent: lipgloss.Color("#b7a8ff"),

		SelectionBg: lipgloss.Color("#f2efd7"),
		SelectionFg: lipgloss.Color("#11131c"),

		Green:  lipgloss.Color("#7ad97a"),
		Yellow: lipgloss.Color("#e0b84f"),
		Red:    lipgloss.Color("#ff7a7a"),
		Blue:   lipgloss.Color("#7aa2ff"),
	}
}
