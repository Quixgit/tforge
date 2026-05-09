package app

import (
	"os"
	"path/filepath"
	"strings"
)

func (m Model) readModuleReadme() string {
	if m.moduleTarget == nil {
		return ""
	}

	candidates := []string{
		"README.md",
		"README",
		"readme.md",
		"docs.md",
	}

	for _, name := range candidates {
		path := filepath.Join(m.moduleTarget.Dir, name)

		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		text := strings.TrimSpace(string(data))
		if text != "" {
			return text
		}
	}

	return ""
}
