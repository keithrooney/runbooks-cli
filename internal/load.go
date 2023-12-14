package internal

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Load(filepath string) (*Runbook, error) {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return nil, fmt.Errorf("Error: file %s does not exist\n", filepath)
	}
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("Error: failed to read file %s\n", filepath)
	}
	var runbook *Runbook
	if err := yaml.Unmarshal(bytes, &runbook); err != nil {
		return nil, fmt.Errorf("Error: failed to parse file %s\n", filepath)
	}
	return runbook, nil
}
