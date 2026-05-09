package app

func (m Model) renderScanError() string {
	msg := errorStyle.Render("Scan failed") + "\n\n" + m.err.Error()

	if requiresTerraformInit(m.err.Error()) {
		msg += "\n\n" + warningStyle.Render("Recovery available:")
		msg += "\n" + dimStyle.Render("Press I to run terraform init, then Ctrl+r to retry.")
	} else {
		msg += "\n\n" + dimStyle.Render("Press Ctrl+r to retry | q to quit")
	}

	return msg
}
