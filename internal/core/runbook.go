package core

import (
	_ "embed"
)

type Runbook struct {
	Title      string     `yaml:"title"`
	Details    string     `yaml:"details"`
	Impact     string     `yaml:"impact"`
	Mitigation Mitigation `yaml:"mitigation"`
}
