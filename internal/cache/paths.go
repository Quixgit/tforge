package cache

import (
	"crypto/sha1"
	"encoding/hex"
	"os"
	"path/filepath"
)

func ProjectDir(projectDir string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	sum := sha1.Sum([]byte(projectDir))
	id := hex.EncodeToString(sum[:])[:16]

	dir := filepath.Join(home, ".tforge", "cache", id)

	if err := os.MkdirAll(dir, 0o700); err != nil {
		return "", err
	}

	return dir, nil
}

func PlanPath(projectDir string) (string, error) {
	dir, err := ProjectDir(projectDir)
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "last.tfplan"), nil
}
