package history

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type Entry struct {
	ID      string    `json:"id"`
	Time    time.Time `json:"time"`
	Dir     string    `json:"dir"`
	Engine  string    `json:"engine"`
	Action  string    `json:"action"`
	Success bool      `json:"success"`
	Logs    []string  `json:"logs"`
}

func Save(entry Entry) error {
	base, err := baseDir()
	if err != nil {
		return err
	}

	dir := filepath.Join(base, "history")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}

	if entry.ID == "" {
		entry.ID = time.Now().Format("20060102-150405")
	}
	if entry.Time.IsZero() {
		entry.Time = time.Now()
	}

	data, err := json.MarshalIndent(entry, "", "  ")
	if err != nil {
		return err
	}

	path := filepath.Join(dir, entry.ID+".json")
	return os.WriteFile(path, data, 0o600)
}

func baseDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".tforge"), nil
}
