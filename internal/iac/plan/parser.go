package plan

import (
	"encoding/json"
	"fmt"
	"os"
)

func ParseFile(path string) (Plan, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Plan{}, err
	}

	return Parse(data)
}

func Parse(data []byte) (Plan, error) {
	var p Plan

	if err := json.Unmarshal(data, &p); err != nil {
		return Plan{}, fmt.Errorf("parse terraform plan json: %w", err)
	}

	return p, nil
}
