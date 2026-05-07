package project

import (
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
)

func Discover(root string) ([]Target, error) {
	targets := []Target{}
	seen := map[string]bool{}

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if d.IsDir() && shouldSkipDir(d.Name()) {
			return filepath.SkipDir
		}

		if d.IsDir() {
			return nil
		}

		dir := filepath.Dir(path)
		base := filepath.Base(path)

		if seen[dir] {
			return nil
		}

		switch base {
		case "terragrunt.hcl":
			targets = append(targets, Target{
				Name: relName(root, dir),
				Dir:  dir,
				Kind: KindTerragrunt,
			})
			seen[dir] = true

		case ".tofu.lock.hcl":
			targets = append(targets, Target{
				Name: relName(root, dir),
				Dir:  dir,
				Kind: KindTofu,
			})
			seen[dir] = true

		default:
			if strings.HasSuffix(base, ".tf") {
				targets = append(targets, Target{
					Name: relName(root, dir),
					Dir:  dir,
					Kind: KindTerraform,
				})
				seen[dir] = true
			}
		}

		return nil
	})

	sort.Slice(targets, func(i, j int) bool {
		if targets[i].Kind == targets[j].Kind {
			return targets[i].Name < targets[j].Name
		}
		return targets[i].Kind < targets[j].Kind
	})

	return targets, err
}

func shouldSkipDir(name string) bool {
	switch name {
	case ".git", ".terraform", ".terragrunt-cache", "node_modules", "vendor":
		return true
	default:
		return false
	}
}

func relName(root, dir string) string {
	rel, err := filepath.Rel(root, dir)
	if err != nil || rel == "." {
		return filepath.Base(dir)
	}

	return rel
}
