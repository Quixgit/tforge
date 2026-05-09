package app

func (m Model) taskRuntime() RuntimeInfo {
	rt := m.runtime

	if m.activeTarget != nil {
		rt.Dir = m.activeTarget.Dir

		if m.activeTarget.Kind != "" {
			rt.Engine = string(m.activeTarget.Kind)
		}
	}

	return rt
}
