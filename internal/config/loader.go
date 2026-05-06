package config

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func Load(dir string) (Config, error) {
	cfg := Default()

	paths := []string{
		filepath.Join(dir, "tforge.yaml"),
		filepath.Join(dir, ".tforge.yaml"),
	}

	home, err := os.UserHomeDir()
	if err == nil {
		paths = append(paths, filepath.Join(home, ".tforge", "config.yaml"))
	}

	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err == nil {
			if err := yaml.Unmarshal(data, &cfg); err != nil {
				return Config{}, err
			}
			return cfg, nil
		}

		if !errors.Is(err, os.ErrNotExist) {
			return Config{}, err
		}
	}

	return cfg, nil
}
