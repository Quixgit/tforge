package moduleparser

import (
	"os"
	"path/filepath"
	"regexp"
)

var (
	variableRe = regexp.MustCompile(`variable\s+"([^"]+)"`)
	outputRe   = regexp.MustCompile(`output\s+"([^"]+)"`)
	resourceRe = regexp.MustCompile(`resource\s+"([^"]+)"\s+"([^"]+)"`)
	providerRe = regexp.MustCompile(`provider\s+"([^"]+)"`)
)

func Parse(dir string) (Module, error) {
	var m Module

	files, err := filepath.Glob(filepath.Join(dir, "*.tf"))
	if err != nil {
		return m, err
	}

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			continue
		}

		content := string(data)

		for _, match := range variableRe.FindAllStringSubmatch(content, -1) {
			m.Variables = append(m.Variables, match[1])
		}

		for _, match := range outputRe.FindAllStringSubmatch(content, -1) {
			m.Outputs = append(m.Outputs, match[1])
		}

		for _, match := range providerRe.FindAllStringSubmatch(content, -1) {
			m.Providers = append(m.Providers, match[1])
		}

		for _, match := range resourceRe.FindAllStringSubmatch(content, -1) {
			m.Resources = append(
				m.Resources,
				match[1]+"."+match[2],
			)
		}
	}

	return m, nil
}
