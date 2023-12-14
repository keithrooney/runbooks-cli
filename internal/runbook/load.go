package runbook

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Load(filepath string) (*Runbook, error) {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return nil, fmt.Errorf("Error: file does not exist\n")
	}
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("Error: unable to read file\n")
	}
	var runbook *Runbook
	if err := yaml.Unmarshal(bytes, &runbook); err != nil {
		return nil, fmt.Errorf("Error: unexpected file format\n")
	}
	return runbook, nil
}
