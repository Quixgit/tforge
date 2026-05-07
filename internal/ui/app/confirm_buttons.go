package app

func confirmButton(label string, focused bool, danger bool) string {
	style := borderStyle.Padding(0, 3)

	if danger {
		style = style.BorderForeground(colorRed).Foreground(colorRed)
	}

	if focused {
		style = cursorStyle.Padding(0, 3)
	}

	return style.Render(label)
}
