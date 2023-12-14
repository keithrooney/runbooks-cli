package runbook

type Step struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
}
