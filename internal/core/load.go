package core

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Load(filepath string) (*Runbook, error) {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file error: %w\n", err)
	}
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("file error: %w\n", err)
	}
	var runbook *Runbook
	if err := yaml.Unmarshal(bytes, &runbook); err != nil {
		return nil, fmt.Errorf("file error: %w\n", err)
	}
	return runbook, nil
}
