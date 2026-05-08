package moduleparser

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	variableBlockRe = regexp.MustCompile(`variable\s+"([^"]+)"\s*\{([^}]*)\}`)
	outputRe        = regexp.MustCompile(`output\s+"([^"]+)"`)
	resourceRe      = regexp.MustCompile(`resource\s+"([^"]+)"\s+"([^"]+)"`)
	providerRe      = regexp.MustCompile(`provider\s+"([^"]+)"`)

	typeRe        = regexp.MustCompile(`type\s*=\s*([^\n]+)`)
	defaultRe     = regexp.MustCompile(`default\s*=\s*([^\n]+)`)
	descriptionRe = regexp.MustCompile(`description\s*=\s*"([^"]+)"`)
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

		parseVariables(&m, content)
		parseOutputs(&m, content)
		parseProviders(&m, content)
		parseResources(&m, content)
	}

	return m, nil
}

func parseVariables(m *Module, content string) {
	for _, match := range variableBlockRe.FindAllStringSubmatch(content, -1) {
		v := Variable{
			Name:     match[1],
			Required: true,
		}

		body := match[2]

		if x := typeRe.FindStringSubmatch(body); len(x) > 1 {
			v.Type = strings.TrimSpace(x[1])
		}

		if x := defaultRe.FindStringSubmatch(body); len(x) > 1 {
			v.Default = strings.TrimSpace(x[1])
			v.Required = false
		}

		if x := descriptionRe.FindStringSubmatch(body); len(x) > 1 {
			v.Description = x[1]
		}

		m.Variables = append(m.Variables, v)
	}
}

func parseOutputs(m *Module, content string) {
	for _, match := range outputRe.FindAllStringSubmatch(content, -1) {
		m.Outputs = append(m.Outputs, Output{
			Name: match[1],
		})
	}
}

func parseProviders(m *Module, content string) {
	for _, match := range providerRe.FindAllStringSubmatch(content, -1) {
		m.Providers = append(m.Providers, match[1])
	}
}

func parseResources(m *Module, content string) {
	for _, match := range resourceRe.FindAllStringSubmatch(content, -1) {
		m.Resources = append(
			m.Resources,
			match[1]+"."+match[2],
		)
	}
}
