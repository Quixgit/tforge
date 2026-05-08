package deps

import (
	"os"
	"path/filepath"
	"regexp"
)

var (
	moduleSourceRe = regexp.MustCompile(`source\s*=\s*"([^"]+)"`)
)

func ParseDependencies(dir string) ([]string, error) {
	var out []string

	files, err := filepath.Glob(filepath.Join(dir, "*.tf"))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			continue
		}

		content := string(data)

		for _, match := range moduleSourceRe.FindAllStringSubmatch(content, -1) {
			if len(match) > 1 {
				out = append(out, match[1])
			}
		}
	}

	return unique(out), nil
}

func unique(in []string) []string {
	seen := map[string]bool{}
	var out []string

	for _, v := range in {
		if seen[v] {
			continue
		}

		seen[v] = true
		out = append(out, v)
	}

	return out
}
