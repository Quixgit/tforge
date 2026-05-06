package history

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
)

func List(limit int) ([]Entry, error) {
	base, err := baseDir()
	if err != nil {
		return nil, err
	}

	dir := filepath.Join(base, "history")

	files, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return []Entry{}, nil
		}
		return nil, err
	}

	entries := make([]Entry, 0, len(files))

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		path := filepath.Join(dir, f.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		var entry Entry
		if err := json.Unmarshal(data, &entry); err != nil {
			continue
		}

		entries = append(entries, entry)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Time.After(entries[j].Time)
	})

	if limit > 0 && len(entries) > limit {
		entries = entries[:limit]
	}

	return entries, nil
}
