package app

func (m Model) renderModuleInspectorOverlay(view string) string {
	return centeredLayer(
		view,
		m.renderModuleInspector(),
		6,
		m.width,
		m.height,
	)
}

func (m Model) renderGraphOverlayOnView(view string) string {
	return centeredLayer(
		view,
		m.renderGraphOverlay(),
		6,
		m.width,
		m.height,
	)
}
